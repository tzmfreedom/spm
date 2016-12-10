package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	log "github.com/Sirupsen/logrus"
	_ "github.com/k0kubun/pp"
	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Username       string
	Password       string
	Endpoint       string
	ApiVersion     string
	PollSeconds    int
	TimeoutSeconds int
	PackageFile    string
}

const (
	ExitCodeOK int = iota
	ExitCodeError
)

const (
	DEFAULT_REPOSITORY        string = "github.com"
)

var client *ForceClient = nil

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
					Name:        "username, u",
					Destination: &config.Username,
					EnvVar:      "SF_USERNAME",
				},
				cli.StringFlag{
					Name:        "password, p",
					Destination: &config.Password,
					EnvVar:      "SF_PASSWORD",
				},
				cli.StringFlag{
					Name:        "endpoint, e",
					Value:       "login.salesforce.com",
					Destination: &config.Endpoint,
					EnvVar:      "SF_ENDPOINT",
				},
				cli.StringFlag{
					Name:        "apiversion",
					Value:       "38.0",
					Destination: &config.ApiVersion,
					EnvVar:      "SF_APIVERSION",
				},
				cli.IntFlag{
					Name:        "pollSeconds",
					Value:       5,
					Destination: &config.PollSeconds,
					EnvVar:      "SF_POLLSECONDS",
				},
				cli.IntFlag{
					Name:        "timeoutSeconds",
					Value:       0,
					Destination: &config.TimeoutSeconds,
					EnvVar:      "SF_TIMEOUTSECONDS",
				},
				cli.StringFlag{
					Name:        "packages, P",
					Destination: &config.PackageFile,
				},
			},
			Action: func(c *cli.Context) error {
				urls := []string{}
				if config.PackageFile != "" {
					packageFile, err := readPackageFile(config.PackageFile)
					if err != nil {
						log.Error(err)
					}
					for _, pkg := range packageFile.Packages {
						urls = append(urls, convertToUrl(pkg))
					}
				} else {
					urls = []string{convertToUrl(c.Args().First())}
				}
				return install(urls, &config)
			},
		},
	}

	statusCode := ExitCodeOK
	err := app.Run(os.Args)
	if err != nil {
		statusCode = ExitCodeError
		log.Error(err)
	}
	os.Exit(statusCode)
}

func install(urls []string, config *Config) error {
	err := checkConfigration(config)
	if err != nil {
		return err
	}
	err = setClient(config)
	if err != nil {
		return err
	}

	for _, url := range urls {
		r := regexp.MustCompile(`^(https://([^/]+?)/([^/]+?)/([^/@]+?))(@([^/]+))?$`)
		group := r.FindAllStringSubmatch(url, -1)
		uri := group[0][1]
		directory := group[0][4]
		branch := group[0][6]
		if branch == "" {
			branch = "master"
		}
		log.Info("Clone repository from " + uri + " (branch: " + branch + ")")

		err = installToSalesforce(uri, directory, branch, config)
		if err != nil {
			return err
		}
	}
	return nil
}

func setClient(config *Config) error {
	client = NewForceClient(config.Endpoint, config.ApiVersion)
	err := client.Login(config.Username, config.Password)
	if err != nil {
		return err
	}
	return nil
}

func convertToUrl(target string) string {
	url := target
	r := regexp.MustCompile(`^[^/]+/[^/@]+(@[^/]+)?$`)
	if r.MatchString(url) {
		url = DEFAULT_REPOSITORY + "/" + url
	}
	return "https://" + url
}

type PackageFile struct {
	Packages []string
}

func readPackageFile(packageFileName string) (*PackageFile, error) {
	packageFile := PackageFile{}
	readBody, err := ioutil.ReadFile(packageFileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(readBody), &packageFile)
	if err != nil {
		return nil, err
	}
	return &packageFile, nil
}

func checkConfigration(config *Config) error {
	if config.Username == "" {
		return errors.New("Username is required")
	}
	if config.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

func installToSalesforce(url string, directory string, branch string, config *Config) error {
	cloneDir := filepath.Join(os.TempDir(), directory)
	err := cloneFromRemoteRepository(cloneDir, url, branch)
	if err != nil {
		return err
	}
	defer cleanTempDirectory(cloneDir)
	err = deployToSalesforce(filepath.Join(cloneDir, "src"), config)
	if err != nil {
		return err
	}
	return nil
}

func cleanTempDirectory(directory string) error {
	if err := os.RemoveAll(directory); err != nil {
		log.Error(err)
		return nil
	}
	return nil
}

func cloneFromRemoteRepository(directory string, url string, paramBranch string) error {
	branch := "master"
	if paramBranch != "" {
		branch = paramBranch
	}
	r, err := git.NewFilesystemRepository(directory)
	err = r.Clone(&git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
	})
	if err != nil {
		return err
	}

	// ... retrieving the branch being pointed by HEAD
	ref, _ := r.Head()

	// ... retrieving the commit object
	commit, _ := r.Commit(ref.Hash())

	// ... we get all the files from the commit
	files, _ := commit.Files()

	_, err = r.Head()
	err = files.ForEach(func(f *git.File) error {
		abs := filepath.Join(directory, "src", f.Name)
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
	if err != nil {
		return err
	}
	return nil
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
				paths = append(paths, fmt.Sprintf(filepath.Join("%s", ""), rel))
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

func zipDirectory(directory string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zwriter := zip.NewWriter(buf)

	files, err := find(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		absPath, _ := filepath.Abs(filepath.Join(directory, file))
		info, _ := os.Stat(absPath)

		f, err := zwriter.Create(filepath.Join("src", file))

		if info.IsDir() {
			continue
		}

		body, err := ioutil.ReadFile(absPath)
		if err != nil {
			return nil, err
		}
		f.Write(body)
	}

	zwriter.Close()
	return buf, nil
}

func deployToSalesforce(directory string, config *Config) error {
	buf, err := zipDirectory(directory)
	if err != nil {
		return err
	}

	response, err := client.Deploy(buf.Bytes())
	if err != nil {
		return err
	}

	err = checkDeployStatus(response.Result.Id, config)
	if err != nil {
		return err
	}
	log.Info("Deploy is successful")

	return nil
}

func checkDeployStatus(resultId *ID, config *Config) error {
	totalTime := 0
	for {
		time.Sleep(time.Duration(config.PollSeconds) * time.Second)
		log.Info("Check Deploy Result...")

		response, err := client.CheckDeployStatus(resultId)
		if err != nil {
			return err
		}
		if response.Result.Done {
			return nil
		}
		if config.TimeoutSeconds != 0 {
			totalTime += config.PollSeconds
			if totalTime > config.TimeoutSeconds {
				log.Error("Deploy is timeout. Please check release status for the deployment")
				return nil
			}
		}
	}
}
