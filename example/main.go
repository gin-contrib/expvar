package main

import (
	"github.com/gin-contrib/expvar"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/debug/vars", expvar.Handler())
	r.Run(":8080")
}
