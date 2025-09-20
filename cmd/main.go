package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/viveksingh-01/jarvis-api/config"
	"github.com/viveksingh-01/jarvis-api/routes"
)

func init() {
	loadEnvVariables()
}

func main() {
	fmt.Println("Welcome to JARVIS API")

	// Initialize Gemini client
	config.InitializeGemini()

	// Instantiate Gorilla Mux for routing
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	// Load port from .env if available
	port := "9090"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

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
