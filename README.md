[![Build Status](https://travis-ci.org/tzmfreedom/spm.svg?branch=master)](https://travis-ci.org/tzmfreedom/spm)

# SPM

Salesforce Package Manager

## Install

For Linux and macOS user
```bash
$ curl -sL http://install.freedom-man.com/spm | bash
```

If you want to install zsh completion, add --zsh-completion option
```bash
$ curl -sL http://install.freedom-man.com/spm | bash -s -- --zsh-completion
```

If you use macOS, home brew is available for installation.
```bash
$ brew tap tzmfreedom/spm
$ brew install spm
```

You can use docker to use spm command.
```bash
$ docker run --rm tzmfree/spm install {REPO} -u {USERNAME] -p {PASSWORD}

```

If you want to use latest version, execute following command.
```bash
$ go get -u github.com/tzmfreedom/spm
```

For Windows user, use Linux virtual machine, such as docker or vagrant.

## Usage

```bash
$ spm [global options] command [command options] [arguments...]

COMMANDS:
     install, i  Install salesforce metadata on public remote repository(i.g. github) or salesforce org
     clone, c    Download metadata from salesforce organization
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Install Package

```bash
$ spm install [command options] [arguments...]

OPTIONS:
   --username value, -u value   [$SF_USERNAME]
   --password value, -p value   [$SF_PASSWORD]
   --endpoint value, -e value  (default: "login.salesforce.com") [$SF_ENDPOINT]
   --apiversion value          (default: "38.0") [$SF_APIVERSION]
   --pollSeconds value         (default: 5) [$SF_POLLSECONDS]
   --timeoutSeconds value      (default: 0) [$SF_TIMEOUTSECONDS]
   --packages value, -P value
```

* Install from remote repository

```bash
$ spm install https://github.com/{USER}/{REPOSITORY} -u {USERNAME} -p {PASSWORD}

example:
$ spm install https://github.com/tzmfreedom/apex_tdclient -u hoge -p fuga
```

You can specify repository with following format.
```
{REMOTE_REPOSITORY_BASE}/{USER}/{REPOSITORY} # i.g. https://github.com/tzmfreedom/apex_tdclient
{REMOTE_REPOSITORY_BASE}/{USER}/{REPOSITORY}/{SUB_DIRECTORY} # i.g. https://github.com/tzmfreedom/apex_tdclient/sample/repositories/dependencies
{USER}/{REPOSITORY} # i.g. tzmfreedom/apex_tdclient
{USER}/{REPOSITORY}/{SUB_DIRECTORY} # i.g. tzmfreedom/spm/sample/repositories/dependencies
```

* Install from salesforce

```bash
$ spm install sf://{username}:{password}@{salesforce endpoint}?path={package file path}&version={version}

example:
$ spm install sf://hoge:fuga@login.salesforce.com?path=./package.toml&version=38.0
```

* Install from package.yml
```bash
$ spm install -u {USERNAME} -p {PASSWORD} -P package.yml
```

Dependency package yaml format

```yaml
packages:
  - tzmfreedom/apex-util1
  - tzmfreedom/apex-util2
  - tzmfreedom/apex-util3
```

Sandbox

```bash
$ spm install https://github.com/{USER}/{REPOSITORY} -u {USERNAME} -p {PASSWORD} -e test.salesforce.com
```

## Download metadata from salesforce

```bash
$ spm clone sf://{username}:{password}@{salesforce endpoint}?path={package file path}&version={version}

example:
$ spm clone sf://hoge:fuga@login.salesforce.com
```

### Package File Format

The package file format for downloading from salesforce is toml.

* example
```toml
version = 37.0

[[types]]
name = "ApexClass"
members = ["Hoge", "Fuga"]

[[types]]
name = "CustomObject"
members = ["Account", "Contact"]

[[types]]
name = "ApexPage"
members = ["HogePage"]

[[types]]
name = "ApexTrigger"
members = ["HogeTrigger"]

[[types]]
name = "Layout"
members = ["Task-Task Layout"]

[[types]]
name = "Profile"
members = ["Admin"]
```

## Contribute

Just send pull request if needed or fill an issue!

## License

The MIT License See [LICENSE](https://github.com/tzmfreedom/spm/blob/master/LICENSE) file.
