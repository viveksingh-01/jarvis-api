package routes

import (
	"github.com/gorilla/mux"
	"github.com/viveksingh-01/jarvis-api/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/conversation", handlers.HandleConversation)
}
