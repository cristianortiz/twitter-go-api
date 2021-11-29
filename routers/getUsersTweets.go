package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cristianortiz/twitter-go-api/bd"
)

//GetUsersTweets is the controller function to process the request for request and paginate
//all the tweets from a specific user
func GetUsersTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "the userID parameter is missing", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "number of pages is missing", http.StatusBadRequest)
		return
	}
	//coverting the string URL param to an integer
	pageTemp, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "pages must be greater than 0", http.StatusBadRequest)
	}
	page := int64(pageTemp)
	response, done := bd.GetUserTweetsDB(ID, page)
	if !done {
		http.Error(w, "Error tring to get tweets", http.StatusBadRequest)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
