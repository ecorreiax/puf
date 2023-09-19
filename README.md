# puf

[![Go Report Card](https://goreportcard.com/badge/github.com/ecorreiax/puf)](https://goreportcard.com/report/github.com/ecorreiax/puf)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ecorreiax/puf)
![GitHub](https://img.shields.io/github/license/ecorreiax/puf)
![GitHub issues](https://img.shields.io/github/issues/ecorreiax/puf)

**PUF** _(Pejorative Username Filter)_ is a package built in Go that uses bloom's filter to validate if a specific username is definitely not present in a dataset that might represent a pejorative username that shouldn't be used.  

## Why use puf

With puf, you can quickly check a username against our hand-picked dataset to make sure it's cool with community rules. That way, we all get to hang out in a safer space. We're always tweaking our dataset to keep up with the latest hacks.

## Install

With a [correctly configured](https://go.dev/doc/install#testing) Go toolchain:

```
go get -u github.com/ecorreiax/puf
```

## Examples

Kick things off by passing a username string to the `Check()` method. This returns a boolean, tipping you off if the username is invalid or not.

```go
func createUser() {
  u := "username"
  invalid := puf.Check(u)
  if invalid {
   // abort creation
  }
  // continue creation
}
```

## Core team

- [@ecorreiax](https://github.com/ecorreiax)


## License

This project is under [MIT License](https://github.com/ecorreiax/puf/blob/main/LICENSE).
