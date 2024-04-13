package handler

import (
	"bta/internal/db"
	"bta/internal/server"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtSecret = []byte("secret_key")

func Register(w http.ResponseWriter, r *http.Request) {
	var creds db.User

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		server.Error(map[string]interface{}{"message": "Data is not provided", "status": 400}, w)
		return
	}

	err = db.InsertUser(creds)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error(), "status": 400}, w)
		return
	}

	server.Ok(map[string]interface{}{"message": "success"}, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds db.User

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		server.Error(map[string]interface{}{"message": "Data is not provided", "status": 400}, w)
		return
	}

	err = db.ValidatePassword(creds.Email, creds.Password)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error(), "status": 400}, w)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": creds.Email,
		"exp":   time.Now().Add(120 * time.Minute).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		server.Error(map[string]interface{}{"message": "Error generating token", "status": 400}, w)
		return
	}

	server.Ok(map[string]interface{}{"message": "Signed-in successful", "token": tokenString, "status": 200}, w)
}
