package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"srcd.works/go-git.v4"
	"srcd.works/go-git.v4/plumbing"
)

const DEFAULT_REPOSITORY string = "github.com"

type Installer interface {
	Install(urls []string) error
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
	urlStack   []string
}

func NewSalesforceInstaller(logger Logger, downloader Downloader, config *config) (*SalesforceInstaller, error) {
	i := &SalesforceInstaller{
		logger:     logger,
		config:     config,
		downloader: downloader,
		urlStack:   []string{},
	}
	err := i.init()
	return i, err
}

func (i *SalesforceInstaller) init() (err error) {
	if i.config.IsCloneOnly {
		return nil
	}
	if i.config.Username == "" {
		return errors.New("Username is required")
	}
	if i.config.Password == "" {
		return errors.New("Password is required")
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

func (i *SalesforceInstaller) Install(uris []string) error {
	for _, uri := range uris {

		err := i.installToSalesforce(uri)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *SalesforceInstaller) installToSalesforce(uri string) error {
	i.addUrlStack(uri)
	defer i.popUrlStack()
	files, err := i.downloader.Download(uri)
	if err != nil {
		return err
	}
	err = i.loadDependencies(uri)

	zc := NewZipConverter()
	if files, err = zc.Convert(files); err != nil {
		return err
	}

	err = i.deployToSalesforce(files[0].Body)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}

func (i *SalesforceInstaller) cloneFromRemoteRepository(directory string, url string, paramBranch string, retry bool) (err error) {
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
		i.logger.Warningf("repository non empty: %s", directory)
		i.logger.Infof("remove directory: %s", directory)
		err = cleanTempDirectory(directory)
		if err != nil {
			return
		}
		err = i.cloneFromRemoteRepository(directory, url, paramBranch, true)
	}
	return
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
	i.logger.Infof("%s: Deploy is successful", i.getTopStack())

	return nil
}

func (i *SalesforceInstaller) checkDeployStatus(resultId *ID) error {
	totalTime := 0
	for {
		time.Sleep(time.Duration(i.config.PollSeconds) * time.Second)
		i.logger.Infof("%s: Check Deploy Result...", i.getTopStack())

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
				i.logger.Errorf("%s: Deploy is timeout. Please check release status for the deployment", i.getTopStack())
				return nil
			}
		}
	}
}

func (i *SalesforceInstaller) loadDependencies(uri string) error {
	_, _, dir, _ := extractInstallParameter(uri)
	targetFile := filepath.Join(dir, "package.yml")
	_, err := os.Stat(targetFile)
	if err != nil {
		return nil
	}
	packageFile, err := readPackageFile(targetFile)
	if err != nil {
		return err
	}
	urls := []string{}
	for _, pkg := range packageFile.Packages {
		url, err := convertToUrl(pkg)
		if err != nil {
			return err
		}
		urls = append(urls, url)
	}
	di, err := NewSalesforceInstaller(i.logger, i.downloader, i.config)
	if err != nil {
		return err
	}
	return di.Install(urls)
}

func (i *SalesforceInstaller) addUrlStack(url string) error {
	i.urlStack = append(i.urlStack, url)
	return nil
}

func (i *SalesforceInstaller) getTopStack() string {
	return i.urlStack[len(i.urlStack)-1]
}

func (i *SalesforceInstaller) popUrlStack() {
	i.urlStack = i.urlStack[:len(i.urlStack)-1]
}
