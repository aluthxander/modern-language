package belajar9web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	// logic web
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
	contentXPoweredBy := request.Header.Get("x-powered-by")
	fmt.Fprint(writer, contentXPoweredBy)
}

func TestRequestHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-powered-by", "belajar golang")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}