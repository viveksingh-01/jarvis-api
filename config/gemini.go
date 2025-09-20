package config

import (
	"context"
	"log"
	"os"

	"github.com/viveksingh-01/jarvis-api/handlers"
	"google.golang.org/genai"
)

func InitializeGemini() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}

	c, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal("Failed to initialize Gemini client:", err)
	}
	handlers.Client = c
	log.Println("Gemini client initialized successfully")
}
