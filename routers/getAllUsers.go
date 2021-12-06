package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cristianortiz/twitter-go-api/bd"
)

//GetAllUsers is thr controller func to list the users from DD
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	p := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")
	//the page parameter must be converted to int
	pagTemp, err := strconv.Atoi(p)
	if err != nil {
		http.Error(w, "page parameter must be an int greater than zero", http.StatusBadRequest)
		return
	}
	//convert the pagTemp int into a int64 type
	page := int64(pagTemp)
	result, status := bd.GetAllUsersDB(userID, page, search, typeUser)
	if !status {
		http.Error(w, "Error trying to get all users", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
