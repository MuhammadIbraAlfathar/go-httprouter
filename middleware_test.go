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

type LogMiddleware struct {
	Handler http.Handler
}

func (m *LogMiddleware) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Middleware")
	m.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	middleware := LogMiddleware{Handler: router}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHttp(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))

}
