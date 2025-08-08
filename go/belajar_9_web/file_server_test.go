package belajar9web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resourses")
	fileServer := http.FileServer(directory)
	
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resourses
var resourses embed.FS
func TestFileServerGolangEmbed(t *testing.T) {
	directory, errDir := fs.Sub(resourses, "resourses")
	if errDir != nil {
		panic(errDir)
	}
	fileServer := http.FileServer(http.FS(directory))
	
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}