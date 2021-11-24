package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

//UpdateUser is a controller to call the same method of db package to update user data
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var t models.User
	//decodes the body JSON from the http request (body is a stream can be procesed only in this line)
	err := json.NewDecoder(r.Body).Decode(&t) //after that line Body disappear from memory
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}
	//userID is one of the global var created via JWT
	status, err := bd.UpdateUserDB(t, userID)
	if err != nil {
		http.Error(w, "An Error occurs:"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "An Error occurs on update user data ", 400)
		return
	}
	//return a status of http object
	w.WriteHeader(http.StatusCreated)
}
