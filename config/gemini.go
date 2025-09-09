package config

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

func InitializeGemini() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	_, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Printf("Failed to initialize Gemini client: %v", err)
	}
}
