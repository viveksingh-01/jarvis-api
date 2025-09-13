package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/viveksingh-01/jarvis-api/models"
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

	var req models.ConversationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "Invalid request body",
		})
	}
}
