package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/viveksingh-01/jarvis-api/models"
	"github.com/viveksingh-01/jarvis-api/utils"
	"google.golang.org/genai"
)

const GEMINI_MODEL = "gemini-2.0-flash"

var (
	Client   *genai.Client
	sessions = make(map[string]*genai.Chat)
	mu       sync.Mutex
	cfg      *genai.GenerateContentConfig
)

func init() {
	systemInstructionText := "You are JARVIS and I'm Tony Stark. Please keep the conversation short, in less than 20 words."
	p := &genai.Part{Text: systemInstructionText}
	cfg = &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{Parts: []*genai.Part{p}},
	}
}

func HandleConversation(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidatePostRequestMethod(w, r) {
		return
	}
	if !utils.ValidateRequestBody(w, r) {
		return
	}

	// Check if client is initialized
	if Client == nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			Error: "Gemini client not initialized",
		})
		return
	}

	var req models.ConversationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	mu.Lock()
	session, exist := sessions[req.Email]
	if !exist {
		session, err := Client.Chats.Create(r.Context(), GEMINI_MODEL, cfg, nil)
		if err != nil {
			log.Println("Error creating new chat session", err.Error())
		}
		sessions[req.Email] = session
	}
	mu.Unlock()

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
