package routers

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func performRequest(router *gin.Engine, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)

	// we will get errors if not set Content-Type to JSON
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	return w
}
