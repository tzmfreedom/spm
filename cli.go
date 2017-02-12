package main

import (
	"errors"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/golang/go/src/regexp"
	"github.com/urfave/cli"
)

type CLI struct {
	installer Installer
	Config    *Config
	logger    *Logger
	Error     error
}

type PackageFile struct {
	Packages []string
}

const (
	APP_VERSION        string = "0.1.0"
	DEFAULT_REPOSITORY string = "github.com"
)

func NewCli() *CLI {
	logger := NewLogger(os.Stdout, os.Stderr)
	c := &CLI{
		installer: NewSalesforceInstaller(&Config{}, logger),
		logger:    logger,
	}
	return c
}

func (c *CLI) Run(args []string) (err error) {
	err = c.installer.Initialize()
	if err != nil {
		return err
	}
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
					Name:        "clone-only",
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
				c.Error = c.installer.Install(urls)
				return nil
			},
		},
	}

	app.Run(args)
	if c.Error != nil {
		c.logger.Error(c.Error)
	}
	return c.Error
}

func (c *CLI) loadInstallUrls(args cli.Args) ([]string, error) {
	urls := []string{}
	if c.Config.PackageFile != "" {
		packageFile, err := c.readPackageFile()
		if err != nil {
			return nil, err
		}
		for _, pkg := range packageFile.Packages {
			url, err := convertToUrl(pkg)
			if err != nil {
				return nil, err
			}
			urls = append(urls, url)
		}
	} else {
		url, err := convertToUrl(args.First())
		if err != nil {
			return nil, err
		}
		urls = []string{url}
	}
	return urls, nil
}

func convertToUrl(target string) (string, error) {
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

func (c *CLI) readPackageFile() (*PackageFile, error) {
	packageFile := PackageFile{}
	readBody, err := ioutil.ReadFile(c.Config.PackageFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(readBody), &packageFile)
	if err != nil {
		return nil, err
	}
	return &packageFile, nil
}
