package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to JARVIS API")

	// Instantiate Gorilla Mux for routing
	r := mux.NewRouter()

	port := "9090"
	log.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
