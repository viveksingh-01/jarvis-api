package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/viveksingh-01/jarvis-api/database"
	"github.com/viveksingh-01/jarvis-api/models"
	"github.com/viveksingh-01/jarvis-api/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidatePostRequestMethod(w, r) {
		return
	}
	if !utils.ValidateRequestBody(w, r) {
		return
	}
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var user models.User
	// Check if the user already exists based on the 'email'
	err := database.UserCollection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
	if err == nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: "The email is already registered.\n Please try logging in.",
		})
		return
	}

	// TODO
}
