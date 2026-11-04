package expvar

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(t testing.TB, r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequestWithContext(t.Context(), method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHandler(t *testing.T) {
	router := gin.New()
	router.GET("/debug/vars", Handler())

	w := performRequest(t, router, "GET", "/debug/vars")
	assert.Equal(t, 200, w.Code)
}
