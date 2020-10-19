# Gitlabctl

[![Latest Release](https://img.shields.io/github/release/thobianchi/gitlabctl.svg?style=flat-square)](https://github.com/thobianchi/gitlabctl/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/thobianchi/gitlabctl)](https://goreportcard.com/report/github.com/thobianchi/gitlabctl)
[![CI](https://github.com/thobianchi/gitlabctl/workflows/goreleaser/badge.svg)](https://github.com/thobianchi/gitlabctl/actions?query=workflow%3Agoreleaser)

## Overview

This command line utility provides some commands to do operations on gitlab like, create repos, copy environment from project to local shell, launch pipeline and perhaps monitor it.

Not a complete coverage of gitlab APIs but only some useful shortcuts, use it on your own risks:   
"Be wary, triumphant pride precipitates a dizzying fall"

## Usage

Set Context to connect to a Gitlab instance

```
gitlabctl config set-context --contextName git --token 'asdasdasdasd' --gitlabURL 'https://gitlab.com'
gitlabctl --help
```

Get Environment of a Project, including parent groups until root

```
gitlabctl projects get-env --id 637
```

Delete temporary files

```
gitlabctl clean
```

## Installation

```
brew tap thobianchi/tap
brew install thobianchi/tap/gitlabctl
```
