# renamer

A simple rename tool.

## Install

```sh
go get -u github.com/devlights/renamer/cmd/renamer
```

## Run

```sh
renamer -p /path/to/directory [-r] [-v] -old Old-Strings -new New-Strings
```

### Example

```sh
renamer -p ~/tmp -r -old ".txt" -new ".bin"
```
