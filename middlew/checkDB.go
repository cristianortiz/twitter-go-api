package middlew

import (
	"net/http"

	"github.com/cristianortiz/twitter-go-api/bd"
)

//CheckDB middleware function, check the DB connection, receive a http func type and must return the same type of func
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	// CheckDB returns an anonymous function, check if the BD connection is still alive
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(w, "DB Connection Lost..", 500)
			return
		}
		//if DB is connected pass the control and parameters to next element in the handler file request
		next.ServeHTTP(w, r)
	}
}
