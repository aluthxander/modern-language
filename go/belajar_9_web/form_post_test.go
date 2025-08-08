package belajar9web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// logic web
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}


	firstName := request.FormValue("first_name")
	lastName := request.FormValue("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("first_name=Udin&last_name=Putra")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/hello", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
