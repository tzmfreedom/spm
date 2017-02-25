package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	FAILURE_PACKAGE_YML_BLANK      = "./test/fixture/blank.yml"
	FAILURE_PACKAGE_YML_REPO_BLANK = "./test/fixture/repo-blank.yml"
	SUCCESS_PACKAGE_TOML           = "./test/fixture/package.toml"
)

func before() (*CLI, *bytes.Buffer, *bytes.Buffer) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := NewCli()
	cli.logger.Reset(outStream, errStream)
	return cli, outStream, errStream
}

func TestInstallSuccess(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install %s -u %s -p %s", os.Getenv("REPOSITORY"), os.Getenv("USERNAME"), os.Getenv("PASSWORD")), " ")
	cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, fmt.Sprintf("Clone repository from https://github.com/%s (branch: %s)", os.Getenv("REPOSITORY"), "master"))
	assert.Contains(t, outString, "Check Deploy Result...")
	assert.Contains(t, outString, "Deploy is successful")
}

func TestInstallSuccessForSubdir(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install %s -u %s -p %s", os.Getenv("REPOSITORY_SUBDIR"), os.Getenv("USERNAME"), os.Getenv("PASSWORD")), " ")
	cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, fmt.Sprintf("Clone repository from https://github.com/"))
	assert.Contains(t, outString, "Check Deploy Result...")
	assert.Contains(t, outString, "Deploy is successful")
}

func TestInstallFailureNoUsername(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install %s -p %s", os.Getenv("REPOSITORY"), os.Getenv("PASSWORD")), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Username is required")
}

func TestInstallFailureNoPassword(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install %s -u %s", os.Getenv("REPOSITORY"), os.Getenv("USERNAME")), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Password is required")
}

func TestInstallFailureNoRepository(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install -u %s -p %s", os.Getenv("USERNAME"), os.Getenv("PASSWORD")), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Repository not specified")
}

func TestInstallFailureNoPackageYML(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install -u %s -p %s -P %s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "NOPACKAGE.yml"), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "open NOPACKAGE.yml: no such file or directory")
}

func TestInstallFailureInvalidCredentials(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install %s -u hoge -p fuga", os.Getenv("REPOSITORY")), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "INVALID_LOGIN: Invalid username, password, security token; or user locked out.")
}

func TestInstallFailurePackageYmlBlank(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install -u %s -p %s -P %s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), FAILURE_PACKAGE_YML_BLANK), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Repository not specified")
}

func TestInstallFailurePackageYmlRepoBlank(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm install -u %s -p %s -P %s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), FAILURE_PACKAGE_YML_REPO_BLANK), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Repository not specified")
}

func TestDownloadSuccess(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm clone -u %s -p %s -P %s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), SUCCESS_PACKAGE_TOML), " ")
	cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Start Retrieve Request...")
	assert.Contains(t, outString, "Check Retrieve Status...")
	assertIsExists(t, "tmp/unpackaged/classes/HelloSpm_Dep.cls")
	assertIsExists(t, "tmp/unpackaged/classes/HelloSpm_Dep.cls-meta.xml")
	assertIsExists(t, "tmp/unpackaged/package.xml")
	assert.Nil(t, os.RemoveAll("./tmp"))
}

func TestDownloadFailureNoUsername(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm clone -p %s -P %s", os.Getenv("PASSWORD"), SUCCESS_PACKAGE_TOML), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Username is required")
}

func TestDownloadFailureNoPassword(t *testing.T) {
	cli, outStream, _ := before()
	args := strings.Split(fmt.Sprintf("spm clone -u %s -P %s", os.Getenv("USERNAME"), SUCCESS_PACKAGE_TOML), " ")
	_ = cli.Run(args)
	outString := outStream.String()
	assert.Contains(t, outString, "Password is required")
}

func assertIsExists(t *testing.T, filename string) {
	_, err := os.Stat(filename)
	assert.Nil(t, err)
}
