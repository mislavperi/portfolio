package main

import (
	"fmt"
	"net/http"

	"github.com/mislavperi/portfolio/handlers"
	"github.com/mislavperi/portfolio/middleware"
)

func main() {
	muxHandler := http.NewServeMux()

	fmt.Println("Assinging handlers")
	muxHandler.Handle("/", handlers.Index())

	loggedMux := middleware.Logger(muxHandler)

	server := http.Server{
		Handler: loggedMux,
		Addr:    ":8080",
	}
	fmt.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
