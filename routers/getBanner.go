package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/cristianortiz/twitter-go-api/bd"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
	//this can be the id on any user to get their avatar
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is mandatory ", http.StatusBadRequest)
		return
	}
	//get the usef progile based on an id
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "the user does not exists ", http.StatusBadRequest)
		return
	}
	opFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "cannot find the user banner ", http.StatusBadRequest)
	}
	//if the avatar img is founded, Copy send to the writer the binary file
	_, err = io.Copy(w, opFile)
	if err != nil {
		http.Error(w, "Error in banner copy", http.StatusBadRequest)

	}
}
