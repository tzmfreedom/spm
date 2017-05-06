package main

import (
	"errors"
	"fmt"
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

var (
	Version  string
	Revision string
)

const (
	DEFAULT_API_VERSION = "38.0"
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
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
	}

	app := cli.NewApp()
	app.Name = "spm"

	app.Usage = "Salesforce Package Manager"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install salesforce metadata on public remote repository(i.g. github) or salesforce org",
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
				uris, err := loadInstallUrls(c.Config.PackageFile, ctx.Args().First())
				if err != nil {
					return err
				}
				if len(uris) == 0 {
					err = errors.New("Repository not specified")
					return err
				}
				for _, uri := range uris {
					downloader, err := dispatchDownloader(c.logger, uri)
					if err != nil {
						return err
					}

					installer, err := NewSalesforceInstaller(c.logger, downloader, c.Config, uri)
					if err != nil {
						return err
					}
					if err = installer.Install(); err != nil {
						return err
					}
				}
				return err
			},
		},
		{
			Name:    "uninstall",
			Aliases: []string{"u"},
			Usage:   "Uninstall salesforce metadata on public remote repository(i.g. github) or salesforce org",
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
			},
			Action: func(ctx *cli.Context) error {
				uris, err := loadInstallUrls(c.Config.PackageFile, ctx.Args().First())
				if err != nil {
					return err
				}
				if len(uris) == 0 {
					err = errors.New("Repository not specified")
					return err
				}
				for _, uri := range uris {
					downloader, err := dispatchDownloader(c.logger, uri)
					if err != nil {
						return err
					}

					installer, err := NewSalesforceInstaller(c.logger, downloader, c.Config, uri)
					if err != nil {
						return err
					}
					if err = installer.Uninstall(); err != nil {
						return err
					}
				}
				return err
			},
		},
		{
			Name:    "clone",
			Aliases: []string{"c"},
			Usage:   "Download metadata from salesforce organization",
			Flags: []cli.Flag{
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
				uri, err := convertToUrl(ctx.Args().First())
				if err != nil {
					return err
				}
				downloader, err := dispatchDownloader(c.logger, uri)
				if err != nil {
					return err
				}
				files, err := downloader.Download()
				if err != nil {
					return err
				}
				if _, ok := downloader.(*SalesforceDownloader); ok {
					err = unzip(files[0].Body, c.Config.Directory)
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
