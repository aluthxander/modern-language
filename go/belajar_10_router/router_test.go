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

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprintln(writer, "Hello, World!")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Fprintln adds a newline, so adjust the expected value
	assert.Equal(t, "Hello, World!\n", string(body))
}

func TestRouterNotFound(t *testing.T) {
	router := httprouter.New()
	request := httptest.NewRequest("GET", "http://localhost:3000/unknown", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestRouterWithParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/hello/:name", func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		name := ps.ByName("name")
		fmt.Fprintf(writer, "Hello, %s!", name)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/hello/Alice", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello, Alice!", string(body))
}
