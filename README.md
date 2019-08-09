# renamer

A simple rename tool.

## Install

```sh
go get -u github.com/devlights/renamer/cmd/renamer
```

## Run

```sh
renamer -path /path/to/directory [-r] [-v] -old Old-Strings -new New-Strings
```

### Example

```sh
renamer -path ~/tmp -r -old ".txt" -new ".bin"
```
