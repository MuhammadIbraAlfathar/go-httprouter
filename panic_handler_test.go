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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Error: ", error)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("upss")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Error: upss", string(body))

}
