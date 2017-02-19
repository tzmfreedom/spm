package main

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"time"

	"github.com/BurntSushi/toml"
)

type Downloader interface {
	Initialize(config *Config) error
	Download() ([]byte, error)
}

type MetaPackageFile struct {
	Version float64 `toml:"version"`
	Types   []*Type `toml:"types"`
}

type Type struct {
	Name    string   `toml:"name"`
	Members []string `toml:"members"`
}

type SalesforceDownloader struct {
	Config *Config
	Client *ForceClient
	logger Logger
}

func NewSalesforceDownloader(logger Logger) *SalesforceDownloader {
	return &SalesforceDownloader{
		logger: logger,
	}
}

func (i *SalesforceDownloader) Initialize(config *Config) (err error) {
	i.Config = config
	if i.Config.Username == "" {
		return errors.New("Username is required")
	}
	if i.Config.Password == "" {
		return errors.New("Password is required")
	}

	err = i.setClient()
	if err != nil {
		return err
	}
	return nil
}

func (i *SalesforceDownloader) setClient() error {
	i.Client = NewForceClient(i.Config.Endpoint, i.Config.ApiVersion)
	err := i.Client.Login(i.Config.Username, i.Config.Password)
	if err != nil {
		return err
	}
	return nil
}

func (i *SalesforceDownloader) Download() (buf []byte, err error) {
	buf, err = ioutil.ReadFile(i.Config.PackageFile)
	if err != nil {
		return nil, err
	}
	packages := &MetaPackageFile{}
	err = toml.Unmarshal(buf, packages)
	if err != nil {
		return nil, err
	}
	i.logger.Info("Start Retrieve Request...")
	r, err := i.Client.Retrieve(createRetrieveRequest(packages))
	for {
		time.Sleep(2 * time.Second)
		i.logger.Info("Check Retrieve Status...")
		ret_res, err := i.Client.CheckRetrieveStatus(r.Result.Id)
		if err != nil {
			return nil, err
		}
		if ret_res.Result.Done {
			buf = make([]byte, len(ret_res.Result.ZipFile))
			_, err = base64.StdEncoding.Decode(buf, ret_res.Result.ZipFile)
			return buf, err
		}
	}
	return
}
