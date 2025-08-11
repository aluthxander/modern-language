package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("receive request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestRouterMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprintln(writer, "Middleware!")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()
	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Fprintln adds a newline, so adjust the expected value
	assert.Equal(t, "Middleware!\n", string(body))
}
