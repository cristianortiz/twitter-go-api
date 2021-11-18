package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

//Register creates a new user in BD, receives the flow control and request params from middleware CheckDB
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	//decodes the body object of the http request (body is a stream can be procesed only in this line)
	err := json.NewDecoder(r.Body).Decode(&t) //after that line Body disappear from memory
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}
	//email is mandatory field, check it here
	if len(t.Email) == 0 {
		http.Error(w, "Email is mandatory ", 400)
		return
	}
	//password validation
	if len(t.Email) < 6 {
		http.Error(w, "Password minimum length is 6 characters ", 400)
		return
	}
	//check if the email is already in use by another user
	_, found, _ := bd.UserExists(t.Email)
	if found {
		http.Error(w, "Email is already in use", 400)
		return
	}
	//creates a new user account
	_, status, err := bd.CreateNewUser(t)
	if err != nil {
		http.Error(w, "An Error occurs at user register "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "user register attempt is failed ", 400)
		return
	}
	//return a status of http object
	w.WriteHeader(http.StatusCreated)
}
