package belajar9web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	// Set the headers to indicate a file download
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "bad request")
		return
	}

	writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file))

	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFile(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}