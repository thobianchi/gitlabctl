# Gitlabctl

This command line utility provides some commands to do operations on gitlab like, create repos, copy environment from project to local shell, launch pipeline and perhaps monitor it.

Not a complete coverage of gitlab APIs but only some useful shortcuts, use it on your own risks:
"Be wary triumphant pride precipitates a dizzying fall"

## Usage

```
gitlabctl config set-context --contextName git --token 'asdasdasdasd' --gitlabURL 'https://gitlab.com'
gitlabctl --help
```

## Installation

```
brew tap thobianchi/tap
brew install thobianchi/tap/gitlabctl
```
