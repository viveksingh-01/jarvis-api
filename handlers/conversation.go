package handlers

import (
	"net/http"

	"github.com/viveksingh-01/jarvis-api/utils"
	"google.golang.org/genai"
)

var Client *genai.Client

func HandleConversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, utils.ErrorResponse{
			Error: "Invalid request method, please use POST method",
		})
	}
	if r.Body == nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "Request body cannot be empty",
		})
		defer r.Body.Close()
	}
}
