package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/viveksingh-01/jarvis-api/database"
	"github.com/viveksingh-01/jarvis-api/models"
	"github.com/viveksingh-01/jarvis-api/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	if err != mongo.ErrNoDocuments {
		log.Println("Database error: " + err.Error())
		utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			Error: "An internal error occurred,\n Please try again.",
		})
		return
	}

	// Generate hashed-password and store as password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Println("Error occurred while hashing password:", err.Error())
		utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			Error: "The request couldn't be processed.\n Please try again after some time.",
		})
		return
	}
	user.Email = req.Email
	user.Name = req.Name
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	if _, err := database.UserCollection.InsertOne(context.TODO(), user); err != nil {
		log.Println("Error occurred while inserting user's record to DB:", err.Error())
		utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			Error: "The request couldn't be processed.\n Please try again after some time.",
		})
		return
	}
	log.Printf("New user registered: %s", user.Email)

	// TODO
}
