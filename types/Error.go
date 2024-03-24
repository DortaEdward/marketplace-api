package types

import (
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}
	type ErrorResponse struct {
		error string `json:"error"`
	}
	ResponseWithJson(w, code, ErrorResponse{
		error: msg,
	})
}
