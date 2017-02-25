package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

func extractInstallParameter(url string) (uri string, dir string, target_dir string, branch string) {
	fmt.Println(url)
	r := regexp.MustCompile(`^(https://([^/]+?)/([^/]+?)/([^/@]+?))(/([^@]+))?(@([^/]+))?$`)
	group := r.FindAllStringSubmatch(url, -1)
	uri = group[0][1]
	dir = group[0][4]
	target_dir = group[0][6]
	branch = group[0][8]
	if branch == "" {
		branch = "master"
	}
	return
}

func loadInstallUrls(packageFile string, targetName string) ([]string, error) {
	urls := []string{}
	if packageFile != "" {
		packageFile, err := readPackageFile(packageFile)
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
		url, err := convertToUrl(targetName)
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
	r := regexp.MustCompile(`^(https://|sf://).+$`)
	if r.MatchString(url) {
		return url, nil
	}
	r = regexp.MustCompile(`^[^/]+?/[^/@]+?(/[^@]+?)?(@[^/]+)?$`)
	if r.MatchString(url) {
		url = fmt.Sprintf("%s/%s", DEFAULT_REPOSITORY, url)
	}
	return fmt.Sprintf("https://%s", url), nil
}

func readPackageFile(packageFileStr string) (*PackageFile, error) {
	packageFile := PackageFile{}
	readBody, err := ioutil.ReadFile(packageFileStr)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(readBody), &packageFile)
	if err != nil {
		return nil, err
	}
	return &packageFile, nil
}

func unzip(buf []byte, dest string) error {
	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return err
	}

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), 0755)
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
