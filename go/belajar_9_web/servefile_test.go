package belajar9web

import (
	_"embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resourses/ok.html")
	} else {
		http.ServeFile(writer, request, "./resourses/notFound.html")
	}
}

func TestServeFileServer(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resourses/ok.html
var resoursesOk string

//go:embed resourses/notFound.html
var resoursesNotFound string

func ServeFileGolangEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resoursesOk)
	} else {
		fmt.Fprint(writer, resoursesNotFound)
	}
}

func TestServeFileGolangEmbed(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileGolangEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}