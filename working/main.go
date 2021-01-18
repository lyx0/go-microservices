package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lyx0/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.HandleFunc()

	http.ListenAndServe(":8080", nil)
}
