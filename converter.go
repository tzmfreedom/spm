package main

import (
	"archive/zip"
	"bytes"
)

type Converter interface {
	Convert([]*File) []*File
}

type ZipConverter struct {
	fileName string
}

type File struct {
	Name string
	Body []byte
}

func NewZipConverter() *ZipConverter {
	return &ZipConverter{}
}

func (c *ZipConverter) Convert(files []*File) ([]*File, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)

	for _, f := range files {
		zf, err := zw.Create(f.Name)
		if err != nil {
			return nil, err
		}
		if _, err = zf.Write(f.Body); err != nil {
			return nil, err
		}
	}
	zw.Close()
	return []*File{{Body: buf.Bytes()}}, nil
}
