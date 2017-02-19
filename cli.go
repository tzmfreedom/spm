package main

import (
	"errors"
	"os"

	"github.com/urfave/cli"
)

type CLI struct {
	installer  Installer
	downloader Downloader
	Config     *Config
	logger     *Logger
	Error      error
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
		installer:  NewSalesforceInstaller(logger),
		downloader: NewSalesforceDownloader(logger),
		logger:     logger,
		Config:     &Config{},
	}
	return c
}

func (c *CLI) Run(args []string) (err error) {
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
					Name:        "directory, d",
					Destination: &c.Config.Directory,
				},
			},
			Action: func(ctx *cli.Context) error {
				urls, err := loadInstallUrls(c.Config.PackageFile, ctx.Args().First())
				if err != nil {
					c.Error = err
					return nil
				}
				if len(urls) == 0 {
					c.Error = errors.New("Repository not specified")
					return nil
				}
				err = c.installer.Initialize(c.Config)
				if err != nil {
					c.Error = err
					return nil
				}
				c.Error = c.installer.Install(urls)
				return nil
			},
		},
		{
			Name:    "clone",
			Aliases: []string{"c"},
			Usage:   "Download metadata from salesforce organization",
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
					Name:        "package, P",
					Destination: &c.Config.PackageFile,
				},
				cli.StringFlag{
					Name:        "directory, d",
					Value:       "tmp",
					Destination: &c.Config.Directory,
				},
			},
			Action: func(ctx *cli.Context) error {
				err = c.downloader.Initialize(c.Config)
				if err != nil {
					c.Error = err
					return nil
				}
				buf, err := c.downloader.Download()
				if err != nil {
					c.Error = err
					return nil
				}
				err = unzip(buf, c.Config.Directory)
				if err != nil {
					c.Error = err
					return nil
				}
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
