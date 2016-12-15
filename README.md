# expvar

[![Build Status](https://travis-ci.org/gin-contrib/expvar.svg)](https://travis-ci.org/gin-contrib/expvar)
[![codecov](https://codecov.io/gh/gin-contrib/expvar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/expvar)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/expvar)](https://goreportcard.com/report/github.com/gin-contrib/expvar)
[![GoDoc](https://godoc.org/github.com/gin-contrib/expvar?status.svg)](https://godoc.org/github.com/gin-contrib/expvar)

A expvar handler for gin framework, [expvar](https://golang.org/pkg/expvar/) provides a standardized interface to public variables.

## Usage

### Start using it

Download and install it:

```sh
$ go get github.com/gin-contrib/expvar
```

Import it in your code:

```go
import "github.com/gin-contrib/expvar"
```

### Canonical example:

```go
package main

import (
	"github.com/gin-contrib/expvar"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()

	router := gin.Default()
	router.GET("/debug/vars", expvar.Handler())
	router.Run(":8080")
}
```
