package routers

import (
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

func CreateRelations(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get(("id"))
	if len(ID) < 1 {
		http.Error(w, "ID parameter is mandatory", http.StatusBadRequest)

	}
	var rel models.Relations
	rel.UserID = userID //the global var set up through JWT
	rel.FollowedUserID = ID

	status, err := bd.CreateRelationsDB(rel)
	if err != nil {
		http.Error(w, "An Errors occurs on followed a user "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "follow routines fails "+err.Error(), http.StatusBadRequest)
		return

	}
	w.WriteHeader(http.StatusCreated)
}
