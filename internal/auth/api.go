package auth

import (
	"backend/internal/entity"
	"backend/pkg/db"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var lR entity.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&lR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := db.Connection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := db.UserCheckCreds(conn, lR.Email, lR.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var rR entity.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&rR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := db.Connection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	registered, err := db.UserCheckRegistered(conn, rR.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if registered {
		http.Error(w, "Username or email already taken", http.StatusConflict)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(rR.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := db.User{Email: rR.Email, BluetoothID: rR.BluetoothID, Password: string(hashed)}
	userID, err := db.UserRegister(conn, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := GenerateToken(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(RegisterResponse{Token: token})
}
