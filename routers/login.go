package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/jwt"
	"github.com/cristianortiz/twitter-go-api/models"
)

//Login() Endpoint route to receive the login http request
func Login(w http.ResponseWriter, r *http.Request) {
	//set the responseWriter to receive and send json objects
	w.Header().Add("content-type", "application/json")

	var t models.User
	//decode the json body response and pass by reference to t models.User var
	err := json.NewDecoder(r.Body).Decode(&t)
	//if there is an error return a err msg, the err fron DB and cod error response
	if err != nil {
		http.Error(w, "The username or password are incorrect"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required ", 400)
		return
	}

	userDoc, exists, _ := bd.UserExists(t.Email)
	if !exists {
		http.Error(w, "The username or password are incorrect ", 400)
		return
	}
	jwtToken, err := jwt.GeneratesJWT(userDoc)
	if err != nil {
		http.Error(w, "An error occurs to create the token "+err.Error(), 400)
	}
	response := models.LoginResponse{
		Token: jwtToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	//the jswToken recorded in a cookie (optional)
	//cookie expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   jwtToken,
		Expires: expirationTime,
	})
}
