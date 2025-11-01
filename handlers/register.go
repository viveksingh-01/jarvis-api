package handlers

import (
	"net/http"

	"github.com/viveksingh-01/jarvis-api/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidatePostRequestMethod(w, r) {
		return
	}
	if !utils.ValidateRequestBody(w, r) {
		return
	}
	// TODO
}
