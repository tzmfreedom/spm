package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

func zipDirectory(directory string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zwriter := zip.NewWriter(buf)
	defer zwriter.Close()

	files, err := find(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		absPath, _ := filepath.Abs(filepath.Join(directory, file))
		info, _ := os.Stat(absPath)

		f, err := zwriter.Create(filepath.Join("src", file))

		if info.IsDir() {
			continue
		}

		body, err := ioutil.ReadFile(absPath)
		if err != nil {
			return nil, err
		}
		f.Write(body)
	}

	return buf, nil
}

func find(targetDir string) ([]string, error) {
	var paths []string
	err := filepath.Walk(targetDir,
		func(path string, info os.FileInfo, err error) error {
			rel, err := filepath.Rel(targetDir, path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				paths = append(paths, fmt.Sprintf(filepath.Join("%s", ""), rel))
				return nil
			}

			paths = append(paths, rel)

			return nil
		})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func cleanTempDirectory(directory string) error {
	if err := os.RemoveAll(directory); err != nil {
		return err
	}
	return nil
}

func extractInstallParameter(url string) (uri string, dir string, target_dir string, branch string) {
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
	r := regexp.MustCompile(`^[^/]+?/[^/@]+?(/[^@]+?)?(@[^/]+)?$`)
	if r.MatchString(url) {
		url = DEFAULT_REPOSITORY + "/" + url
	}
	return "https://" + url, nil
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
