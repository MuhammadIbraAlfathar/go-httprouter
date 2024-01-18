package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Terjadi kesalahan")
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Terjadi kesalahan", string(body))

}
