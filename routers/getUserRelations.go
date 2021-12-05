package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

//GetUserRelations is the controller function to get from DB the user relations status
func GetUserRelations(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get(("id"))

	var rel models.Relations
	rel.UserID = userID //the global var set up through JWT
	rel.FollowedUserID = ID
	var relStatus models.GetUserRelationsStatus
	status, err := bd.GetUserRelationsDB(rel)
	if err != nil || !status {
		relStatus.Status = false
	} else {
		relStatus.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(relStatus)

}
