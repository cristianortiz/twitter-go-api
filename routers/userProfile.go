package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
)

//UserProfile call db.SearchProfile to get the data of a logged user
func UserProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID param must be send it", http.StatusBadRequest)
		return
	}
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "There is an error in query execution in DB "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
