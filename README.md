# expvar

[![Run Tests](https://github.com/gin-contrib/expvar/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/expvar/actions/workflows/go.yml)
[![Trivy Security Scan](https://github.com/gin-contrib/expvar/actions/workflows/trivy-scan.yml/badge.svg)](https://github.com/gin-contrib/expvar/actions/workflows/trivy-scan.yml)
[![codecov](https://codecov.io/gh/gin-contrib/expvar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/expvar)
[![Go Reference](https://pkg.go.dev/badge/github.com/gin-contrib/expvar.svg)](https://pkg.go.dev/github.com/gin-contrib/expvar)

A expvar handler for gin framework, [expvar](https://golang.org/pkg/expvar/) provides a standardized interface to public variables.

## Usage

### Start using it

Download and install it:

```sh
go get github.com/gin-contrib/expvar
```

Import it in your code:

```go
import "github.com/gin-contrib/expvar"
```

### Canonical example

```go
package main

import (
  "log"

  "github.com/gin-contrib/expvar"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/debug/vars", expvar.Handler())

  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
