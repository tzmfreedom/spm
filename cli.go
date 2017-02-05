package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	_ "github.com/k0kubun/pp"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"srcd.works/go-git.v4"
	"srcd.works/go-git.v4/plumbing"
)

type CLI struct {
	Client *ForceClient
	Config *Config
	Logger *Logger
	Error  error
}

type Config struct {
	Username       string
	Password       string
	Endpoint       string
	ApiVersion     string
	PollSeconds    int
	TimeoutSeconds int
	PackageFile    string
	IsCloneOnly    bool
	Directory      string
}

type PackageFile struct {
	Packages []string
}

const (
	APP_VERSION        string = "0.1.0"
	DEFAULT_REPOSITORY string = "github.com"
)

func (c *CLI) Run(args []string) (err error) {
	if c.Logger == nil {
		c.Logger = NewLogger(os.Stdout, os.Stderr)
	}
	c.Config = &Config{}

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
					Destination: &c.Config.Username,
					EnvVar:      "SF_USERNAME",
				},
				cli.StringFlag{
					Name:        "password, p",
					Destination: &c.Config.Password,
					EnvVar:      "SF_PASSWORD",
				},
				cli.StringFlag{
					Name:        "endpoint, e",
					Value:       "login.salesforce.com",
					Destination: &c.Config.Endpoint,
					EnvVar:      "SF_ENDPOINT",
				},
				cli.StringFlag{
					Name:        "apiversion",
					Value:       "38.0",
					Destination: &c.Config.ApiVersion,
					EnvVar:      "SF_APIVERSION",
				},
				cli.IntFlag{
					Name:        "pollSeconds",
					Value:       5,
					Destination: &c.Config.PollSeconds,
					EnvVar:      "SF_POLLSECONDS",
				},
				cli.IntFlag{
					Name:        "timeoutSeconds",
					Value:       0,
					Destination: &c.Config.TimeoutSeconds,
					EnvVar:      "SF_TIMEOUTSECONDS",
				},
				cli.StringFlag{
					Name:        "packages, P",
					Destination: &c.Config.PackageFile,
				},
				cli.BoolFlag{
					Name:	     "clone-only",
					Destination: &c.Config.IsCloneOnly,
				},
				cli.StringFlag{
					Name:        "directory, -d",
					Destination: &c.Config.Directory,
				},
			},
			Action: func(ctx *cli.Context) error {
				urls, err := c.loadInstallUrls(ctx.Args())
				if err != nil {
					c.Error = err
					return nil
				}
				if len(urls) == 0 {
					c.Error = errors.New("Repository not specified")
					return nil
				}
				c.Error = c.install(urls)
				return nil
			},
		},
	}

	app.Run(args)
	if c.Error != nil {
		c.Logger.Error(c.Error)
	}
	return c.Error
}

func (c *CLI) loadInstallUrls(args cli.Args) ([]string, error) {
	urls := []string{}
	if c.Config.PackageFile != "" {
		packageFile, err := c.readPackageFile(c.Config.PackageFile)
		if err != nil {
			return nil, err
		}
		for _, pkg := range packageFile.Packages {
			url, err := c.convertToUrl(pkg)
			if err != nil {
				return nil, err
			}
			urls = append(urls, url)
		}
	} else {
		url, err := c.convertToUrl(args.First())
		if err != nil {
			return nil, err
		}
		urls = []string{url}
	}
	return urls, nil
}

// Todo: separete CLI class and Installer class.
func (c *CLI) initialize() error {
	err := c.checkConfigration()
	if err != nil {
		return err
	}
	if !c.Config.IsCloneOnly {
		err = c.setClient()
	}
	if err != nil {
		return err
	}
	if c.Config.Directory == "" {
		if c.Config.IsCloneOnly {
			dir, err := os.Getwd()
			if err != nil{
				return err
			}
			c.Config.Directory = dir
		} else {
			c.Config.Directory = os.TempDir()
		}
	}
	return nil
}

func (c *CLI) install(urls []string) error {
	err := c.initialize()
	if err != nil {
		return err
	}
	for _, url := range urls {
		r := regexp.MustCompile(`^(https://([^/]+?)/([^/]+?)/([^/@]+?))(/([^@]+))?(@([^/]+))?$`)
		group := r.FindAllStringSubmatch(url, -1)
		uri := group[0][1]
		directory := group[0][4]
		targetDirectory := group[0][6]
		branch := group[0][8]
		if branch == "" {
			branch = "master"
		}

		err = c.installToSalesforce(uri, directory, targetDirectory, branch)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CLI) setClient() error {
	c.Client = NewForceClient(c.Config.Endpoint, c.Config.ApiVersion)
	err := c.Client.Login(c.Config.Username, c.Config.Password)
	if err != nil {
		return err
	}
	return nil
}

func (c *CLI) convertToUrl(target string) (string, error) {
	if target == "" {
		return "", errors.New("Repository not specified")
	}
	url := target
	r := regexp.MustCompile(`^[^/]+?/[^/@]+?(/[^@]+?)?(@[^/]+)?$`)
	if r.MatchString(url) {
		url = DEFAULT_REPOSITORY + "/" + url
	}
	return "https://" + url, nil
}

func (c *CLI) readPackageFile(configFile string) (*PackageFile, error) {
	packageFile := PackageFile{}
	readBody, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(readBody), &packageFile)
	if err != nil {
		return nil, err
	}
	return &packageFile, nil
}

func (c *CLI) checkConfigration() error {
	if c.Config.IsCloneOnly {
		return nil
	}
	if c.Config.Username == "" {
		return errors.New("Username is required")
	}
	if c.Config.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

func (c *CLI) installToSalesforce(url string, directory string, targetDirectory string, branch string) error {
	cloneDir := filepath.Join(c.Config.Directory, directory)
	c.Logger.Info("Clone repository from " + url + " (branch: " + branch + ")")
	err := c.cloneFromRemoteRepository(cloneDir, url, branch, false)
	if err != nil {
		return err
	}
	if c.Config.IsCloneOnly {
		return nil
	}
	defer c.cleanTempDirectory(cloneDir)
	c.loadDependencies(cloneDir)
	err = c.deployToSalesforce(filepath.Join(cloneDir, targetDirectory))
	if err != nil {
		return err
	}
	return nil
}

func (c *CLI) cleanTempDirectory(directory string) error {
	if err := os.RemoveAll(directory); err != nil {
		return err
	}
	return nil
}

func (c *CLI) loadDependencies(cloneDir string) error {
	targetFile := filepath.Join(cloneDir, "package.yml")
	_, err := os.Stat(targetFile)
	if err != nil {
		return nil
	}
	urls := []string{}
	packageFile, err := c.readPackageFile(targetFile)
	if err != nil {
		return err
	}
	for _, pkg := range packageFile.Packages {
		url, err := c.convertToUrl(pkg)
		if err != nil {
			c.Error = err
			return nil
		}
		urls = append(urls, url)
	}
	return c.install(urls)
}

func (c *CLI) cloneFromRemoteRepository(directory string, url string, paramBranch string, retry bool) (err error) {
	branch := "master"
	if paramBranch != "" {
		branch = paramBranch
	}
	_, err = git.PlainClone(directory, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		SingleBranch:  true,
	})
	if err != nil {
		if err.Error() != "repository already exists" {
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
		err = c.cloneFromRemoteRepository(directory, url, paramBranch, true)
	}
	return
}

func (c *CLI) find(targetDir string) ([]string, error) {
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

func (c *CLI) zipDirectory(directory string) (*bytes.Buffer, error) {
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

func (c *CLI) deployToSalesforce(directory string) error {
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

func (c *CLI) checkDeployStatus(resultId *ID) error {
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
