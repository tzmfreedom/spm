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

	_ "github.com/k0kubun/pp"
	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/yaml.v2"
)

type CLI struct {
	Client *ForceClient
	Config *Config
	Logger *Logger
}

type Config struct {
	Username       string
	Password       string
	Endpoint       string
	ApiVersion     string
	PollSeconds    int
	TimeoutSeconds int
	PackageFile    string
}

type PackageFile struct {
	Packages []string
}

const (
	APP_VERSION        string = "0.1.0"
	DEFAULT_REPOSITORY string = "github.com"
)

func (cl *CLI) Run(args []string) (err error) {
	if cl.Logger == nil {
		cl.Logger = NewLogger(os.Stdout)
	}
	cl.Config = &Config{}

	app := cli.NewApp()
	app.Name = "spm"

	app.Usage = "Salesforce Package Manager"
	app.Version = APP_VERSION
	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install salesforce packages on public remote repository(i.g. github)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "username, u",
					Destination: &cl.Config.Username,
					EnvVar:      "SF_USERNAME",
				},
				cli.StringFlag{
					Name:        "password, p",
					Destination: &cl.Config.Password,
					EnvVar:      "SF_PASSWORD",
				},
				cli.StringFlag{
					Name:        "endpoint, e",
					Value:       "login.salesforce.com",
					Destination: &cl.Config.Endpoint,
					EnvVar:      "SF_ENDPOINT",
				},
				cli.StringFlag{
					Name:        "apiversion",
					Value:       "38.0",
					Destination: &cl.Config.ApiVersion,
					EnvVar:      "SF_APIVERSION",
				},
				cli.IntFlag{
					Name:        "pollSeconds",
					Value:       5,
					Destination: &cl.Config.PollSeconds,
					EnvVar:      "SF_POLLSECONDS",
				},
				cli.IntFlag{
					Name:        "timeoutSeconds",
					Value:       0,
					Destination: &cl.Config.TimeoutSeconds,
					EnvVar:      "SF_TIMEOUTSECONDS",
				},
				cli.StringFlag{
					Name:        "packages, P",
					Destination: &cl.Config.PackageFile,
				},
			},
			Action: func(c *cli.Context) error {
				urls := []string{}
				if cl.Config.PackageFile != "" {
					packageFile, err := cl.readPackageFile(cl.Config.PackageFile)
					if err != nil {
						return nil
					}
					for _, pkg := range packageFile.Packages {
						urls = append(urls, cl.convertToUrl(pkg))
					}
				} else {
					urls = []string{cl.convertToUrl(c.Args().First())}
				}
				err = cl.install(urls)
				return nil
			},
		},
	}

	app.Run(args)
	if err != nil {
		cl.Logger.Error(err)
	}
	return err
}

func (c *CLI) install(urls []string) error {
	err := c.checkConfigration()
	if err != nil {
		return err
	}
	err = c.setClient()
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

		err = c.installToSalesforce(uri, directory, branch)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CLI)setClient() error {
	c.Client = NewForceClient(c.Config.Endpoint, c.Config.ApiVersion)
	err := c.Client.Login(c.Config.Username, c.Config.Password)
	if err != nil {
		return err
	}
	return nil
}

func (c *CLI)convertToUrl(target string) string {
	url := target
	r := regexp.MustCompile(`^[^/]+/[^/@]+(@[^/]+)?$`)
	if r.MatchString(url) {
		url = DEFAULT_REPOSITORY + "/" + url
	}
	return "https://" + url
}

func (c *CLI)readPackageFile(packageFileName string) (*PackageFile, error) {
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

func (c *CLI)checkConfigration() error {
	if c.Config.Username == "" {
		return errors.New("Username is required")
	}
	if c.Config.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

func (c *CLI)installToSalesforce(url string, directory string, branch string) error {
	cloneDir := filepath.Join(os.TempDir(), directory)
	c.Logger.Info("Clone repository from " + url + " (branch: " + branch + ")")
	err := c.cloneFromRemoteRepository(cloneDir, url, branch)
	if err != nil {
		return err
	}
	defer c.cleanTempDirectory(cloneDir)
	err = c.deployToSalesforce(filepath.Join(cloneDir, "src"))
	if err != nil {
		return err
	}
	return nil
}

func (c *CLI)cleanTempDirectory(directory string) error {
	if err := os.RemoveAll(directory); err != nil {
		return err
	}
	return nil
}

func (c *CLI)createFilesystemRepository(directory string, url string, paramBranch string, retry bool) (r *git.Repository, err error){
	branch := "master"
	if paramBranch != "" {
		branch = paramBranch
	}
	r, err = git.NewFilesystemRepository(directory)
	err = r.Clone(&git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
	})
	if err != nil {
		if err.Error() != "repository non empty" {
			return
		}
		if retry == true {
			return
		}
		c.Logger.Warningf("repository non empty: %s", directory)
		c.Logger.Infof("remove directory: %s", directory)
		err = c.cleanTempDirectory(directory)
		if err != nil {
			return
		}
		r, err = c.createFilesystemRepository(directory, url, paramBranch, true)
	}
	return
}

func (c *CLI)cloneFromRemoteRepository(directory string, url string, branch string) error {
	r, err := c.createFilesystemRepository(directory, url, branch, false)
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

func (c *CLI)find(targetDir string) ([]string, error) {
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

func (c *CLI)zipDirectory(directory string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zwriter := zip.NewWriter(buf)
	defer zwriter.Close()

	files, err := c.find(directory)
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

	return buf, nil
}

func (c *CLI)deployToSalesforce(directory string) error {
	buf, err := c.zipDirectory(directory)
	if err != nil {
		return err
	}

	response, err := c.Client.Deploy(buf.Bytes())
	if err != nil {
		return err
	}

	err = c.checkDeployStatus(response.Result.Id)
	if err != nil {
		return err
	}
	c.Logger.Info("Deploy is successful")

	return nil
}

func (c *CLI)checkDeployStatus(resultId *ID) error {
	totalTime := 0
	for {
		time.Sleep(time.Duration(c.Config.PollSeconds) * time.Second)
		c.Logger.Info("Check Deploy Result...")

		response, err := c.Client.CheckDeployStatus(resultId)
		if err != nil {
			return err
		}
		if response.Result.Done {
			return nil
		}
		if c.Config.TimeoutSeconds != 0 {
			totalTime += c.Config.PollSeconds
			if totalTime > c.Config.TimeoutSeconds {
				c.Logger.Error("Deploy is timeout. Please check release status for the deployment")
				return nil
			}
		}
	}
}
