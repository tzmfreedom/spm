package main

import (
	"errors"
	"os"

	"github.com/urfave/cli"
)

type CLI struct {
	Config *config
	logger Logger
}

type PackageFile struct {
	Packages []string
}

const (
	APP_VERSION string = "0.1.1"
)

func NewCli() *CLI {
	logger := NewSpmLogger(os.Stdout, os.Stderr)
	c := &CLI{
		logger: logger,
		Config: &config{},
	}
	cli.OsExiter = func(c int) {}
	cli.ErrWriter = &NullWriter{}
	return c
}

func (c *CLI) Run(args []string) error {
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
				downloader, _ := NewGitDownloader(c.logger, &gitConfig{})
				installer, err := NewSalesforceInstaller(c.logger, downloader, c.Config)
				if err != nil {
					return err
				}
				urls, err := loadInstallUrls(c.Config.PackageFile, ctx.Args().First())
				if err != nil {
					return err
				}
				if len(urls) == 0 {
					err = errors.New("Repository not specified")
					return err
				}
				err = installer.Install(urls)
				return err
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
					Value:       "./package.toml",
					Destination: &c.Config.PackageFile,
				},
				cli.StringFlag{
					Name:        "directory, d",
					Value:       "tmp",
					Destination: &c.Config.Directory,
				},
			},
			Action: func(ctx *cli.Context) error {
				downloader, err := NewSalesforceDownloader(c.logger, c.Config)
				if err != nil {
					return err
				}
				files, err := downloader.Download(c.Config.PackageFile)
				if err != nil {
					return err
				}
				err = unzip(files[0].Body, c.Config.Directory)
				if err != nil {
					return err
				}
				return err
			},
		},
	}

	err := app.Run(args)
	if err != nil {
		c.logger.Error(err)
	}
	return err
}
