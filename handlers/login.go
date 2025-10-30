package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/viveksingh-01/jarvis-api/database"
	"github.com/viveksingh-01/jarvis-api/models"
	"github.com/viveksingh-01/jarvis-api/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidatePostRequestMethod(w, r) {
		return
	}
	if !utils.ValidateRequestBody(w, r) {
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var user models.User
	// Check if the user exists
	err := database.UserCollection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		log.Println("User doesn't exist")
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "User doesn't exist, please create an account.",
		})
		return
	}

	// TODO
}
