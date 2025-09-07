package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to JARVIS API")

	port := "9090"
	log.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
