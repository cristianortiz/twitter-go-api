package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/cristianortiz/twitter-go-api/middlew"
	"github.com/cristianortiz/twitter-go-api/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//ServerHandler manage server setup and requests
func ServerHandler() {
	//captures http content and manage request cod responses
	router := mux.NewRouter()
	//route Register receibe POST request and use CheckDB middleware
	//if CheckDB pass routerr.Register is executed
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	//get environment variable who is define server port connection
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//set permission to access API from everywhere, handling the mux with cors
	handler := cors.AllowAll().Handler(router)
	//server is listen in PORT and cors handle the requests
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}