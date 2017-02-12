package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/go/src/archive/zip"
	"github.com/golang/go/src/bytes"
	"github.com/golang/go/src/path/filepath"
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
