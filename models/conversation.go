package models

type ConversationRequest struct {
	Message string `json:"message"`
	Email   string `json:"email"`
	System  string `json:"system,omitempty"`
}

type ConversationResponse struct {
	Response string `json:"response"`
}
