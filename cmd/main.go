package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	loadEnvVariables()
}

func main() {
	fmt.Println("Welcome to JARVIS API")

	// Instantiate Gorilla Mux for routing
	r := mux.NewRouter()

	port := "9090"
	log.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func loadEnvVariables() {
	if os.Getenv("PRODUCTION") == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading the .env file.")
		}
	}
}
