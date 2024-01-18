package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
