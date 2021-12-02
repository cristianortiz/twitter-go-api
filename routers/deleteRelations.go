package routers

import (
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

//DeleteRelations is the controller function to drop a relation calling the function in DB
func DeleteRelations(w http.ResponseWriter, r *http.Request) {

	var rel models.Relations

	ID := r.URL.Query().Get("id")
	rel.FollowedUserID = ID
	rel.UserID = userID
	status, err := bd.DeleteRelationsDB(rel)
	if err != nil {
		http.Error(w, "An Errors occurs on drop the relation  "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "The relations cannot be droped", http.StatusBadRequest)
		return

	}
	w.WriteHeader(http.StatusCreated)
}
