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
	FAILURE_PACKAGE_YML_BLANK      = "./test/blank.yml"
	FAILURE_PACKAGE_YML_REPO_BLANK = "./test/repo-blank.yml"
)

func before() (*CLI, *bytes.Buffer, *bytes.Buffer) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{logger: NewLogger(outStream, errStream)}
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
