package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/viveksingh-01/jarvis-api/models"
	"github.com/viveksingh-01/jarvis-api/utils"
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

	// TODO
}
