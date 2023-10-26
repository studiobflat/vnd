# vnd (Vietnamese Dev Tools)

Template for REST API server made with Golang.

Features:
- Graceful stop
- Close connections before stop
- Response template error/success/success with pagination

## Installation

To install `vnd` package, you need to install Go.

1. You first need [Go](https://golang.org/) installed then you can use the below Go command to install `vnd`.

```sh
go get -u github.com/thienhaole92/vnd
```

2. Import it in your code:

```go
import "github.com/thienhaole92/vnd"
```

## Quick start

### Starting HTTP server

The example how to use `vnd` with to start REST api in [example](./example/example) folder.

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/runner"
)

func main() {
	options := []runner.RunnerOption{
		runner.BuildMonitorServerOption(runner.DefaultMonitorEchoHook),
		runner.BuildRestServerOption(restServiceHook),
	}
	r := runner.NewRunner(options...)
	r.Run()
}

func restServiceHook(rn *runner.Runner, e *echo.Echo, eg *echo.Group) error {
	log := logger.GetLogger("restServiceHook")
	defer log.Sync()

	return nil
}
```
