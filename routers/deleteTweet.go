package routers

import (
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
)

//DeleteTweet is the controller function to call DeleteTweetDB and actually delete
// a tweet by their id and the userID
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	//ID is the tweet from the http request to be deleted
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id is mandatory ", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTweetDB(ID, userID)
	if err != nil {
		http.Error(w, "An error occurs at delete the tweet"+err.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
