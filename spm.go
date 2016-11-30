package main

import (
  "os"
  "os/user"
  "path/filepath"
  "fmt"
  "io"
  "io/ioutil"
  "bytes"
  "archive/zip"
  "regexp"
  "encoding/base64"
  "log"
  "time"

  _ "github.com/k0kubun/pp"
  "github.com/urfave/cli"
  "gopkg.in/src-d/go-git.v4"
)

type Config struct {
  Username string
  Password string
  Endpoint string
  ApiVersion string
}

const (
  DEFAULT_REPOSITORY string = "github.com"
  DEFAULT_SPMDIRECTORY_NAME string = ".spm"
)

func main() {
  config := Config{}

  app := cli.NewApp()
  app.Name = "spm"

  app.Usage = "salesforce package manager"
  app.Commands = []cli.Command{
    {
      Name:    "install",
      Aliases: []string{"i"},
      Usage:   "install salesforce package",
      Flags: []cli.Flag{
        cli.StringFlag{
          Name: "username, u",
          Destination: &config.Username,
          EnvVar: "SF_USERNAME",
        },
        cli.StringFlag{
          Name: "password, p",
          Destination: &config.Password,
          EnvVar: "SF_PASSWORD",
        },
        cli.StringFlag{
          Name: "endpoint, e",
          Value: "login.salesforce.com",
          Destination: &config.Endpoint,
          EnvVar: "SF_ENDPOINT",
        },
        cli.StringFlag{
          Name: "apiversion",
          Value: "38.0",
          Destination: &config.ApiVersion,
          EnvVar: "SF_APIVERSION",
        },
      },
      Action:  func(c *cli.Context) error {
        url := c.Args().First()
        r := regexp.MustCompile(`^[^/]+/[^/]+$`)
        if r.MatchString(url) {
          url = DEFAULT_REPOSITORY + "/" + url
        }
        log.Println("Clone repository from " + url)

        install("https://" + url, config)
        return nil
      },
    },
    {
      Name:    "init",
      Usage:   "initialize",
      Action:  func(c *cli.Context) error {
        spmDir := getSpmDirectory()
        if err := os.Mkdir(spmDir, 0777); err != nil {
          fmt.Println(err)
        }
        return nil
      },
    },
  }

  app.Run(os.Args)
}

func install(url string, config Config) {
  r := regexp.MustCompile(`^(.+?)/(.+?)/(.+?)$`)
  group := r.FindAllStringSubmatch(url, -1)
  directory := group[0][3]

  cloneDir := getSpmDirectory() + "/" + directory
  cloneFromRemoteRepository(cloneDir, url)
  deployToSalesforce(cloneDir + "/src", config)
  cleanTempDirectory(cloneDir)
}

func getSpmDirectory() (string) {
  usr, _ := user.Current()
  return usr.HomeDir + "/" + DEFAULT_SPMDIRECTORY_NAME
}

func cleanTempDirectory(directory string) (error) {
  if err := os.RemoveAll(directory); err != nil {
    return err
  }
  return nil
}

func cloneFromRemoteRepository(directory string, url string) {
  r, err := git.NewFilesystemRepository(directory)
  err = r.Clone(&git.CloneOptions{
    URL: url,
  })
  if err != nil {
    panic(err)
  }

  // ... retrieving the branch being pointed by HEAD
  ref, _ := r.Head()

  // ... retrieving the commit object
  commit, _ := r.Commit(ref.Hash())

  // ... we get all the files from the commit
  files, _ := commit.Files()

  _, err = r.Head()
  err = files.ForEach(func(f *git.File) error {
    abs := filepath.Join(directory + "/src", f.Name)
    dir := filepath.Dir(abs)

    os.MkdirAll(dir, 0777)
    file, err := os.Create(abs)
    if err != nil {
      return err
    }

    defer file.Close()
    r, err := f.Reader()
    if err != nil {
      return err
    }

    defer r.Close()

    if err := file.Chmod(f.Mode); err != nil {
      return err
    }

    _, err = io.Copy(file, r)
    return err
  })
}

func find(targetDir string) ([]string, error) {

  var paths []string
  err := filepath.Walk(targetDir,
    func(path string, info os.FileInfo, err error) error {
      rel, err := filepath.Rel(targetDir, path)
      if err != nil {
        return err
      }

      if info.IsDir() {
        paths = append(paths, fmt.Sprintf("%s/", rel))
        return nil
      }

      paths = append(paths, rel)

      return nil
    })

  if err != nil {
    return nil, err
  }

  return paths, nil
}

func deployToSalesforce(directory string, config Config) (error) {
  buf := new(bytes.Buffer)
  zwriter := zip.NewWriter(buf)

  files, err := find(directory)
  if err != nil {
    return err
  }

  for _, file := range files {
    absPath, _ := filepath.Abs(directory + "/" + file)
    info, _ := os.Stat(absPath)

    f, err := zwriter.Create("src/" + file)

    if info.IsDir() {
      continue
    }

    body, err := ioutil.ReadFile(absPath)
    if err != nil {
      panic(err)
    }
    f.Write(body)
  }

  zwriter.Close()

  portType := NewMetadataPortType("https://" + config.Endpoint + "/services/Soap/u/" + config.ApiVersion, true, nil)
  loginRequest := LoginRequest{Username: config.Username, Password: config.Password }
  loginResponse, _ := portType.Login(&loginRequest)
  result := loginResponse.LoginResult

  request := Deploy{
    ZipFile: base64.StdEncoding.EncodeToString(buf.Bytes()),
    DeployOptions: nil,
  }
  sessionHeader := SessionHeader{
    SessionId: result.SessionId,
  }
  portType.SetHeader(&sessionHeader)
  portType.SetServerUrl(result.MetadataServerUrl)

  response, err := portType.Deploy(&request)
  if err != nil {
    panic(err)
  }
  log.Println("Deploying...")
  for {
    time.Sleep(5000 * time.Millisecond)
    log.Println("Check Deploy Result...")
    check_request := CheckDeployStatus{AsyncProcessId: response.Result.Id, IncludeDetails: true}
    check_response, err := portType.CheckDeployStatus(&check_request)
    if err != nil {
      panic(err)
    }
    if check_response.Result.Done {
      log.Println("Deploy is successful")
      break
    }
  }

  return nil
}
