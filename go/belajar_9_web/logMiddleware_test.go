package belajar9web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler 
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("before Exceute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("after Exceute Handler")
}

func TestLogMiddleware(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer, "Hello Middleware")
    })

    logMiddleware := &LogMiddleware{
        Handler: mux,
    }

    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()

    logMiddleware.ServeHTTP(rec, req)

    body := rec.Body.String()

    if !strings.Contains(body, "Hello Middleware") {
        t.Errorf("expected body to contain 'Hello Middleware', got %s", body)
    }
}