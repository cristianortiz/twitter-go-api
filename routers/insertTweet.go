package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

//InsertTweet is the controller function to process the request for create a new tweet
//and call the function to insert in DB
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	//Tweet type to store only the message of tweet from responseWriter
	var msg models.Tweet
	//decodes the body JSON from the http request (body is a stream can be procesed only in this line)
	err := json.NewDecoder(r.Body).Decode(&msg) //after that line Body disappear from memory
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
	}
	record := models.TweetsModel{
		//userID is one of the global var created via JWT
		UserID:     userID,
		Message:    msg.Message,
		Created_at: time.Now(),
	}

	_, status, err := bd.InsertTweetDB(record)
	if err != nil {
		http.Error(w, "An Errors occurs on tweet insert "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Tweet cannot be created ", 400)
		return

	}
	w.WriteHeader(http.StatusCreated)
}
