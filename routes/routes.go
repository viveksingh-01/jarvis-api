package routes

import (
	"github.com/gorilla/mux"
	"github.com/viveksingh-01/jarvis-api/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", handlers.Register)
	router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/conversation", handlers.HandleConversation)
}
