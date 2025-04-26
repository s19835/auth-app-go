package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s19835/auth-app-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// decode the json file
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// ensure already existing account
	if _, exists := models.Users[user.Username]; exists {
		http.Error(w, "user already exist", http.StatusBadRequest)
		return
	}

	// hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// save the user
	models.Users[user.Username] = string(hashPassword)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User register successfully✨"))
}
