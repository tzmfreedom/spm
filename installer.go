package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/k0kubun/pp"
	"srcd.works/go-git.v4"
	"srcd.works/go-git.v4/plumbing"
)

type Installer interface {
	Initialize(config *Config) error
	Install(urls []string) error
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

type SalesforceInstaller struct {
	Config *Config
	Client *ForceClient
	logger *Logger
	urlStack []string
}

func NewSalesforceInstaller(logger *Logger) *SalesforceInstaller {
	return &SalesforceInstaller{
		logger: logger,
		urlStack: []string{},
	}
}

func (i *SalesforceInstaller) Initialize(config *Config) (err error) {
	i.Config = config
	if i.Config.IsCloneOnly {
		return nil
	}
	if i.Config.Username == "" {
		return errors.New("Username is required")
	}
	if i.Config.Password == "" {
		return errors.New("Password is required")
	}

	if !i.Config.IsCloneOnly {
		err = i.setClient()
	}
	if err != nil {
		return err
	}
	if i.Config.Directory == "" {
		if i.Config.IsCloneOnly {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			i.Config.Directory = dir
		} else {
			i.Config.Directory = os.TempDir()
		}
	}
	return nil
}

func (i *SalesforceInstaller) setClient() error {
	i.Client = NewForceClient(i.Config.Endpoint, i.Config.ApiVersion)
	err := i.Client.Login(i.Config.Username, i.Config.Password)
	if err != nil {
		return err
	}
	return nil
}

func (i *SalesforceInstaller) Install(urls []string) error {
	for _, url := range urls {
		uri, dir, target_dir, branch := extractInstallParameter(url)

		err := i.installToSalesforce(uri, dir, target_dir, branch)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *SalesforceInstaller) installToSalesforce(url string, directory string, targetDirectory string, branch string) error {
	cloneDir := filepath.Join(i.Config.Directory, directory)
	i.addUrlStack(url)
	defer i.popUrlStack()
	i.logger.Info("Clone repository from " + url + " (branch: " + branch + ")")
	err := i.cloneFromRemoteRepository(cloneDir, url, branch, false)
	if err != nil {
		return err
	}
	if i.Config.IsCloneOnly {
		return nil
	}
	defer cleanTempDirectory(cloneDir)
	targetDirAbsPath := filepath.Join(cloneDir, targetDirectory)
	err = i.loadDependencies(targetDirAbsPath)
	if err != nil {
		return err
	}
	err = i.deployToSalesforce(targetDirAbsPath)
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

func (i *SalesforceInstaller) deployToSalesforce(directory string) error {
	buf, err := zipDirectory(directory)
	if err != nil {
		return err
	}

	response, err := i.Client.Deploy(buf.Bytes())
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
		time.Sleep(time.Duration(i.Config.PollSeconds) * time.Second)
		i.logger.Infof("%s: Check Deploy Result...", i.getTopStack())

		response, err := i.Client.CheckDeployStatus(resultId)
		if err != nil {
			return err
		}
		if response.Result.Done {
			return nil
		}
		if i.Config.TimeoutSeconds != 0 {
			totalTime += i.Config.PollSeconds
			if totalTime > i.Config.TimeoutSeconds {
				i.logger.Errorf("%s: Deploy is timeout. Please check release status for the deployment", i.getTopStack())
				return nil
			}
		}
	}
}

func (i *SalesforceInstaller) loadDependencies(cloneDir string) error {
	targetFile := filepath.Join(cloneDir, "package.yml")
	_, err := os.Stat(targetFile)
	if err != nil {
		return nil
	}
	urls := []string{}
	packageFile, err := readPackageFile(targetFile)
	if err != nil {
		return err
	}
	for _, pkg := range packageFile.Packages {
		url, err := convertToUrl(pkg)
		if err != nil {
			return err
		}
		urls = append(urls, url)
	}
	return i.Install(urls)
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