# SPM

Salesforce Package Manager

## Install

Download binary file and copy it to executable path.

If you want to use latest version, execute following command.

```bash
$ go get github.com/tzmfreedom/spm
```

## Usage

Install Package

```bash
$ spm install github.com/{USER}/{REPOSITORY} -u {USERNAME} -p {PASSWORD}
```

```bash
$ spm install -u {USERNAME} -p {PASSWORD} -P package.yml
```

Package.yml format

```yaml
packages:
  - github.com/tzmfreedom/apex-util1
  - github.com/tzmfreedom/apex-util2
  - github.com/tzmfreedom/apex-util3
```

Sandbox

```bash
$ spm install github.com/{USER}/{REPOSITORY} -u {USERNAME} -p {PASSWORD} -e test.salesforce.com
```