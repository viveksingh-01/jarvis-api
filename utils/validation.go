package utils

import (
	"net/http"
)

// ValidateRequestMethod validates that the request method is POST
func ValidatePostRequestMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		SendErrorResponse(w, http.StatusMethodNotAllowed, ErrorResponse{
			Error: "Invalid request method, please use POST method.",
		})
		return false
	}
	return true
}

// ValidateRequestBody validates that the request body is not nil
func ValidateRequestBody(w http.ResponseWriter, r *http.Request) bool {
	if r.Body == nil {
		SendErrorResponse(w, http.StatusBadRequest, ErrorResponse{
			Error: "Request body cannot be empty",
		})
		defer r.Body.Close()
		return false
	}
	return true
}
