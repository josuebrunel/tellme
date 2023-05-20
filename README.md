[![test](https://github.com/josuebrunel/tellme/actions/workflows/test.yml/badge.svg)](https://github.com/josuebrunel/tellme/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/josuebrunel/tellme/badge.svg?branch=main)](https://coveralls.io/github/josuebrunel/tellme?branch=main)

# tellme

It's a small Go package exposing running app informations.
It's inspired by [go-actuator](https://github.com/sinhashubham95/go-actuator)

NB: I was forced to expose the same information for a Go microservice :D

## Installation

```bash
go get github.com/josuebrunel/tellme
```

## How to use

```go

import "github.com/josuebrunel/tellme"

// These variables are expected to be set by the LDFLAGS arguments
// LDFLAGS would be set at compile time by the CI/CD pipeline of any code
// which leverages this library

var (
    EnvName         string // Environment name
    AppName         string // Application/Service name
    AppVersion      string // Application/Service version
    CommitAuthor    string // CommitAuthor - The username/email of the person who authored the commit
    CommitID        string // CommitID - The SHA1 checksum of the commit
    CommitTime      string // CommitTime - The time that the commit occurred
    BuildTime       string // BuildTime - Timestamp that the build occurred
    Branch          string // Branch - The branch the commit exists in
)

// Initialize your app

app := tellme.NewApp(
    EnvName,
    AppName,
    AppVersion,
    BuildTime,
    CommitAuthor,
    CommitID,
    CommitTime,
    Branch,
)

// Get Info
app.GetInfo() // returns a json serializable struct
// Get Env
app.GetEnv() // returns map[string]string
// Get Metrics
app.GetMetrics() // returns a json serializable struct
// Get Threaddump
app.GetThreadDump() // returns []byte
```
