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

func TestRouterWithNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/hello/:name/items/:itemid", func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		name := ps.ByName("name")
		itemId := ps.ByName("itemid")
		text := "Product "+ name +"! Item ID: " + itemId 
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/hello/Alice/items/123", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product Alice! Item ID: 123", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/image/*image", func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		image := ps.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/image/small/abc.jpg", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image : /small/abc.jpg", string(body))
}
