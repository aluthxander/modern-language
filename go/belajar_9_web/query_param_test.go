package belajar9web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writter http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writter, "Hello World")
	}else{
		fmt.Fprintf(writter, "Hello %s", name)
	}
}

func TestSayHello(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()
	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleParam(writter http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	names  := query["name"]
	fmt.Fprint(writter, strings.Join(names, " "))
}

func TestMultipleParam(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Udin&name=Doni", nil)
	recorder := httptest.NewRecorder()
	MultipleParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

