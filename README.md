[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=VBXHBYFU44T5W&source=url)
[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/virtomize/uii-go-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/virtomize/uii-go-cli)](https://goreportcard.com/report/github.com/virtomize/uii-go-cli)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/virtomize/uii-go-cli/blob/master/LICENSE)
[![Website](https://img.shields.io/badge/%40-Virtomize-%231e828c)](https://virtomize.com)
[![Twitter](https://badgen.net/badge/icon/twitter?icon=twitter&label)](https://twitter.com/virtomize)
[![LinkedIn](https://img.shields.io/badge/linkedIn-%20-blue.svg?style=social&logo=linkedin)](https://www.linkedin.com/company/virtomize/)

# UII CLI client
This repository contains a command line client for [**Virtomize Unattended Install Images API**](https://uii.virtomize.com/). 
It is intended both for day to day use, and to demonstrate the usage of the API for other projects. 

## Requirements
In order to use this client, you need an API token form the cloud api.
A token can be created using the [Virtomize website](https://uii.virtomize.com/).
This token can be passed to each  command using `--token=123456`.
Alternatively, the client will use the value provided in the `UIIAPITOKEN` environment variable.

## Usage
The client supports the following commands:

### Help
Use `--help` or `-h` to print a list of all available commands.

### Building an ISO

The `build` command can be used to build an ISO.
The following example command will build an ISO file named `foo.iso`, which will install Debian linux in version 9 using `hostname` as the hostname. 

Example: 
```shell
uii-go-cli build foo.iso debian 11 hostname --token=1234567890ABCDEFGHIJKL
```

### List available operation systems

The `ls os` command can be used to list all available operating systems.

Example:
```shell
uii-go-cli ls os
```

### List available packages for an operating systems

The `ls package` command can be used to list all available operating systems.

Example:
```shell 
uii-go-cli ls package debian 11
```

This command can produce quite a bit of output.
Piping the result into less is a good way to address this. 

```shell
uii-go-cli ls package debian 11 | less
```

# Contribution

Thank you for participating to this project.
Please see our [Contribution Guidlines](https://github.com/virtomize/uii-go-cli/blob/master/CONTRIBUTING.md) for more information.

## Pre-Commit

This repo uses [pre-commit hooks](https://pre-commit.com/). Please install pre-commit and do `pre-commit install`

## Conventional Commits

Format commit messaged according to [Conventional Commits standard](https://www.conventionalcommits.org/en/v1.0.0/).

## Semantic Versioning

Whenever you need to version something make use of [Semantic Versioning](https://semver.org).
