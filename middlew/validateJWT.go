package middlew

import (
	"net/http"

	"github.com/cristianortiz/twitter-go-api/routers"
)

//ValidateJWT middleware function checks and validate the JWT from http request
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//returns if the token send as parameter is a valid one
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
			return
		}
		//if there is no error in token next pass the control
		next.ServeHTTP(w, r)
	}
}
