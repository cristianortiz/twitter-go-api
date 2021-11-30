package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	//get the file element from request
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "banner parameter is mandatory "+err.Error(), http.StatusBadRequest)
		return
	}
	//split the file, to extract the extension as string[1] is the file extension
	var ext = strings.Split(handler.Filename, ".")[1]
	//config the upload filename to the banner and renamed with the userID
	var renamed string = "upload/banners/" + userID + "." + ext
	//open file destination and config  open mod and permissions
	f, err := os.OpenFile(renamed, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error to upload banner "+err.Error(), http.StatusBadRequest)
		return
	}
	//copy the file from http response in  destination f
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error to upload banner "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	user.Banner = userID + "." + ext
	status, err := bd.UpdateUserDB(user, userID)
	if err != nil || !status {
		http.Error(w, "Error to create banner in DB "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusCreated))

}
