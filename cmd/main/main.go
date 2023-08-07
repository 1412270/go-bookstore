package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1412270/go-bookstore/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
