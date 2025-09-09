package handlers

import (
	"net/http"

	"google.golang.org/genai"
)

var Client *genai.Client

func HandleConversation(w http.ResponseWriter, r *http.Request) {
	// TODO
}
