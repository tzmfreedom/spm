package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const DEFAULT_REPOSITORY string = "github.com"

type Installer interface {
	Install() error
}

type timeoutError struct {
	error
}

type config struct {
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

type SalesforceInstaller struct {
	config     *config
	client     *ForceClient
	downloader Downloader
	logger     Logger
	uri        string
}

func NewSalesforceInstaller(logger Logger, downloader Downloader, config *config, uri string) (*SalesforceInstaller, error) {
	i := &SalesforceInstaller{
		logger:     logger,
		config:     config,
		downloader: downloader,
		uri:        uri,
	}
	err := i.init()
	return i, err
}

func (i *SalesforceInstaller) init() (err error) {
	if i.config.IsCloneOnly {
		return nil
	}
	if i.config.Username == "" {
		return errors.New("[Installer] Username is required")
	}
	if i.config.Password == "" {
		return errors.New("[Installer] Password is required")
	}

	if !i.config.IsCloneOnly {
		err = i.setClient()
	}
	if err != nil {
		return err
	}
	if i.config.Directory == "" {
		if i.config.IsCloneOnly {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			i.config.Directory = dir
		} else {
			i.config.Directory = os.TempDir()
		}
	}
	return nil
}

func (i *SalesforceInstaller) setClient() error {
	i.client = NewForceClient(i.config.Endpoint, i.config.ApiVersion)
	err := i.client.Login(i.config.Username, i.config.Password)
	if err != nil {
		return err
	}
	return nil
}

func (i *SalesforceInstaller) Install() error {
	files, err := i.downloader.Download()
	if err != nil {
		return err
	}
	if err = i.loadDependencies(i.uri, files); err != nil {
		return err
	}

	if _, ok := i.downloader.(*GitDownloader); ok {
		zc := NewZipConverter()
		files, err = zc.Convert(files)
		if err != nil {
			return err
		}
	}

	err = i.deployToSalesforce(files[0].Body)
	if err != nil {
		if _, ok := err.(timeoutError); ok {
			errors.New(fmt.Sprintf("%s: Deploy is timeout. Please check release status for the deployment", i.uri))
		}
		return err
	}
	i.logger.Infof("%s: Deploy is successful", i.uri)
	if err != nil {
		return err
	}

	return nil
}

func (i *SalesforceInstaller) deployToSalesforce(bytes []byte) error {
	response, err := i.client.Deploy(bytes)

	if err != nil {
		return err
	}

	err = i.checkDeployStatus(response.Result.Id)
	if err != nil {
		return err
	}

	return nil
}

func (i *SalesforceInstaller) checkDeployStatus(resultId *ID) error {
	totalTime := 0
	for {
		time.Sleep(time.Duration(i.config.PollSeconds) * time.Second)
		i.logger.Infof("%s: Check Deploy Result...", i.uri)

		response, err := i.client.CheckDeployStatus(resultId)
		if err != nil {
			return err
		}
		if response.Result.Done {
			return nil
		}
		if i.config.TimeoutSeconds != 0 {
			totalTime += i.config.PollSeconds
			if totalTime > i.config.TimeoutSeconds {
				return timeoutError{}
			}
		}
	}
}

func (i *SalesforceInstaller) loadDependencies(uri string, files []*File) error {
	_, _, dir, _, err := extractInstallParameter(uri)
	if err != nil {
		return nil
	}
	targetFile := filepath.Join(dir, "package.yml")
	exists := false
	for _, file := range files {
		if file.Name == targetFile {
			exists = true
			break
		}
	}
	if exists == false {
		return nil
	}
	packageFile, err := readPackageFile(targetFile)
	if err != nil {
		return err
	}
	for _, pkg := range packageFile.Packages {
		uri, err := convertToUrl(pkg)
		if err != nil {
			return err
		}
		downloader, err := dispatchDownloader(i.logger, uri)
		if err != nil {
			return err
		}
		di, err := NewSalesforceInstaller(i.logger, downloader, i.config, uri)
		if err != nil {
			return err
		}
		di.Install()
	}
	return nil
}
