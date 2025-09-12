package models

type ConversationRequest struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}

type ConversationResponse struct {
	Response string `json:"response"`
}
