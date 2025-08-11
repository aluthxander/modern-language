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

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Method Not Allowed")
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprintln(writer, "This is a POST request")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Fprintln adds a newline, so adjust the expected value
	assert.Equal(t, "Method Not Allowed", string(body))
}