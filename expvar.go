package expvar

import (
	"expvar"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Handler for gin framework
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		c.Header("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte("{\n"))
		first := true
		expvar.Do(func(kv expvar.KeyValue) {
			if !first {
				_, _ = w.Write([]byte(",\n"))
			}
			first = false
			fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
		})
		_, _ = w.Write([]byte("\n}\n"))
		c.AbortWithStatus(200)
	}
}
