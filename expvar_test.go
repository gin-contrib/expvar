package expvar

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHandler(t *testing.T) {
	router := gin.New()
	router.GET("/debug/vars", Handler())

	w := performRequest(router, "GET", "/debug/vars")
	assert.Equal(t, w.Code, 200)
}
