package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/viveksingh-01/jarvis-api/models"
	"github.com/viveksingh-01/jarvis-api/utils"
	"google.golang.org/genai"
)

const GEMINI_MODEL = "gemini-2.0-flash"

var (
	Client   *genai.Client
	sessions = make(map[string]*genai.Chat)
)

func HandleConversation(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidatePostRequestMethod(w, r) {
		return
	}
	if !utils.ValidateRequestBody(w, r) {
		return
	}

	var req models.ConversationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	session := sessions[req.Email]

	response, err := session.SendMessage(r.Context(), genai.Part{Text: req.Message})
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ConversationResponse{Response: response.Text()})
}
