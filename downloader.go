package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"srcd.works/go-git.v4"
	"srcd.works/go-git.v4/plumbing"
	"srcd.works/go-git.v4/plumbing/object"
	"srcd.works/go-git.v4/storage/memory"

	"bytes"

	"github.com/BurntSushi/toml"
)

type gitConfig struct {
}

type Downloader interface {
	Download(string) ([]*File, error)
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
	config *config
	client *ForceClient
	logger Logger
}

func NewSalesforceDownloader(logger Logger, config *config) (*SalesforceDownloader, error) {
	d := &SalesforceDownloader{
		logger: logger,
		config: config,
	}
	err := d.init()
	return d, err
}

func (i *SalesforceDownloader) init() (err error) {
	if i.config.Username == "" {
		return errors.New("Username is required")
	}
	if i.config.Password == "" {
		return errors.New("Password is required")
	}

	err = i.setClient()
	if err != nil {
		return err
	}
	return nil
}

func (i *SalesforceDownloader) setClient() error {
	i.client = NewForceClient(i.config.Endpoint, i.config.ApiVersion)
	err := i.client.Login(i.config.Username, i.config.Password)
	if err != nil {
		return err
	}
	return nil
}

func (i *SalesforceDownloader) Download(path string) ([]*File, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	packages := &MetaPackageFile{}
	err = toml.Unmarshal(buf, packages)
	if err != nil {
		return nil, err
	}
	i.logger.Info("Start Retrieve Request...")
	r, err := i.client.Retrieve(createRetrieveRequest(packages))
	if err != nil {
		return nil, err
	}
	for {
		time.Sleep(2 * time.Second)
		i.logger.Info("Check Retrieve Status...")
		ret_res, err := i.client.CheckRetrieveStatus(r.Result.Id)
		if err != nil {
			return nil, err
		}
		if ret_res.Result.Done {
			zb := make([]byte, len(ret_res.Result.ZipFile))
			_, err = base64.StdEncoding.Decode(zb, ret_res.Result.ZipFile)
			return []*File{&File{Body: zb}}, err
		}
	}
	return nil, nil // Todo: error handling
}

type GitDownloader struct {
	logger Logger
	config *gitConfig
}

func NewGitDownloader(logger Logger, config *gitConfig) (*GitDownloader, error) {
	return &GitDownloader{
		logger: logger,
		config: config,
	}, nil
}

func (d *GitDownloader) Download(uri string) ([]*File, error) {
	uri, _, _, branch := extractInstallParameter(uri)
	d.logger.Infof("Clone repository from %s (branch: %s)", uri, branch)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		SingleBranch:  true,
		URL:           uri,
	})
	if err != nil {
		return nil, err
	}

	ref, _ := r.Head()
	commit, _ := r.Commit(ref.Hash())

	gfiles, err := commit.Files()
	files := make([]*File, 0)
	err = gfiles.ForEach(func(f *object.File) error {
		reader, err := f.Reader()
		if err != nil {
			return err
		}

		b := new(bytes.Buffer)
		if _, err := b.ReadFrom(reader); err != nil {
			return err
		}
		files = append(files, &File{Name: f.Name, Body: b.Bytes()})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
