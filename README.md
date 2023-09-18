# puf

[![Go Report Card](https://goreportcard.com/badge/github.com/ecorreiax/puf)](https://goreportcard.com/report/github.com/ecorreiax/puf)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ecorreiax/puf)
![GitHub](https://img.shields.io/github/license/ecorreiax/puf)
![GitHub issues](https://img.shields.io/github/issues/ecorreiax/puf)

**puf** (Pejorative username filter) is a package built in Go that uses bloom's filter to validate if a specific username is definitely not present in a dataset that might represent a pejorative username that shouldn't be used.  

## What's a Bloom Filter

> A Bloom filter is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set".
>
> <cite>[Source from Wikipedia][1]</cite>

[1]: https://en.wikipedia.org/wiki/Bloom_filter

## Why use puf

With puf, you can quickly check a username against our hand-picked dataset to make sure it's cool with community rules. That way, we all get to hang out in a safer space. We're always tweaking our dataset to keep up with the latest tricks from the bad guys.

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
