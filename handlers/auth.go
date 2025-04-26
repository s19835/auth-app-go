package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s19835/auth-app-go/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
}
