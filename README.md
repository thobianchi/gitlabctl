# Gitlabctl

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

Or from current directory

```
gitlabctl projects get-env
```

Delete temporary files

```
gitlabctl clean
```

## Installation

```
pip install --user .
```
