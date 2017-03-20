package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"github.com/BurntSushi/toml"
)

type gitConfig struct {
	uri string
}
type salesforceConfig struct {
	username    string
	password    string
	endpoint    string
	packagePath string
	apiVersion  string
}

type Downloader interface {
	Download() ([]*File, error)
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
	config *salesforceConfig
	client *ForceClient
	logger Logger
}

func NewSalesforceDownloader(logger Logger, config *salesforceConfig) (*SalesforceDownloader, error) {
	d := &SalesforceDownloader{
		logger: logger,
		config: config,
	}
	err := d.init()
	return d, err
}

func (d *SalesforceDownloader) init() (err error) {
	if d.config.username == "" {
		return errors.New("[Downloader] Username is required")
	}
	if d.config.password == "" {
		return errors.New("[Downloader] Password is required")
	}

	err = d.setClient()
	if err != nil {
		return err
	}
	return nil
}

func (d *SalesforceDownloader) setClient() error {
	d.client = NewForceClient(d.config.endpoint, d.config.apiVersion)
	err := d.client.Login(d.config.username, d.config.password)
	if err != nil {
		return err
	}
	return nil
}

func (d *SalesforceDownloader) Download() ([]*File, error) {
	buf, err := ioutil.ReadFile(d.config.packagePath)
	if err != nil {
		return nil, err
	}
	packages := &MetaPackageFile{}
	err = toml.Unmarshal(buf, packages)
	if err != nil {
		return nil, err
	}
	d.logger.Info("Start Retrieve Request...")
	r, err := d.client.Retrieve(createRetrieveRequest(packages))
	if err != nil {
		return nil, err
	}
	for {
		time.Sleep(2 * time.Second)
		d.logger.Info("Check Retrieve Status...")
		ret_res, err := d.client.CheckRetrieveStatus(r.Result.Id)
		if err != nil {
			return nil, err
		}
		if ret_res.Result.Done {
			zb := make([]byte, len(ret_res.Result.ZipFile))
			_, err = base64.StdEncoding.Decode(zb, ret_res.Result.ZipFile)
			return []*File{{Body: zb}}, err
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

func (d *GitDownloader) Download() ([]*File, error) {
	uri, _, _, branch, err := extractInstallParameter(d.config.uri)
	if err != nil {
		return nil, err
	}
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
		files = append(files, &File{Name: filepath.Join("unpackaged", f.Name), Body: b.Bytes()})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func dispatchDownloader(logger Logger, uri string) (Downloader, error) {
	r := regexp.MustCompile(`^https://([^/]+?)/([^/]+?)/([^/@]+?)(/([^@]+))?(\?([^/]+))?$`)
	if r.MatchString(uri) {
		return NewGitDownloader(logger, &gitConfig{uri: uri})
	}
	r = regexp.MustCompile(`^sf://([^/]*?):([^/]*)@([^/]+?)(\?.+)?$`)
	if r.MatchString(uri) {
		group := r.FindAllStringSubmatch(uri, -1)
		path, version, err := parseQuery(group[0][4])
		if err != nil {
			return nil, err
		}
		return NewSalesforceDownloader(logger, &salesforceConfig{
			username:    group[0][1],
			password:    group[0][2],
			endpoint:    group[0][3],
			packagePath: path,
			apiVersion:  version,
		})
	}
	return nil, errors.New("Invalid downloader")
}

func parseQuery(q string) (path string, version string, err error) {
	m, err := url.ParseQuery(q)
	if err != nil {
		return
	}
	if len(m["path"]) == 0 || m["path"][0] == "" {
		path = "./package.toml"
	} else {
		path = m["path"][0]
	}
	if len(m["version"]) == 0 || m["version"][0] == "" {
		version = DEFAULT_API_VERSION
	} else {
		version = m["version"][0]
	}
	return
}
