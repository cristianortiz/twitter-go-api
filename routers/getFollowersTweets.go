package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cristianortiz/twitter-go-api/bd"
)

//GetFollowersTweets is the controller function to get followers tweets from DB
func GetFollowersTweets(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "page parameter is mandatory", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page parameter must be greater than 0", http.StatusBadRequest)
		return
	}

	response, correct := bd.GetFollowersTweetsDB(userID, page)
	if !correct {
		http.Error(w, "Error getting the tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json ")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
