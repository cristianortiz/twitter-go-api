package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "avatar parameter is mandatory "+err.Error(), http.StatusBadRequest)
		return
	}
	//split the file, to extract the extension as string
	var ext = strings.Split(handler.Filename, ".")[1]
	//config the upload filename to the avatar and renamed with the userID
	var renamed string = "upload/avatars/" + userID + "." + ext
	//open file destination and config  open mod and permissions
	f, err := os.OpenFile(renamed, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error to upload img "+err.Error(), http.StatusBadRequest)
		return
	}
	//copy the file from http response in  destination f
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error to upload img "+err.Error(), http.StatusBadRequest)
		return

	}

	var user models.User
	user.Avatar = userID + "." + ext
	status, err := bd.UpdateUserDB(user, userID)
	if err != nil || !status {
		http.Error(w, "Error to create avatar in DB "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusCreated))
}
