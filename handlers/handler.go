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

//ServerHandler manage server setup,  requests and defines routes
func ServerHandler() {
	//captures http content and manage request cod responses
	router := mux.NewRouter()
	//route Register receibe POST request and use CheckDB middleware

	//----------ROUTES--------------------------------------------------
	//if CheckDB pass routerr.Register is executed
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	//uodate user data endpoint
	router.HandleFunc("/update", middlew.CheckDB(middlew.ValidateJWT(routers.UpdateUser))).Methods("PUT")
	//login an existing user and return the jwt for authentication
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	//inside checkDB midd put another one to verify the JWT
	router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.UserProfile))).Methods("GET")
	//inside checkDB midd put another one to verify the JWT
	router.HandleFunc("/createtweet", middlew.CheckDB(middlew.ValidateJWT(routers.InsertTweet))).Methods("POST")
	//get all the tweets paginated from an user
	router.HandleFunc("/gettweets", middlew.CheckDB(middlew.ValidateJWT(routers.GetUsersTweets))).Methods("GET")
	//delete a specific tweet of an user
	router.HandleFunc("/deletetweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")
	//upload an avatar image for an user
	router.HandleFunc("/uploadavatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")

	router.HandleFunc("/uploadbanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getbanner", middlew.CheckDB(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")
	router.HandleFunc("/getavatar", middlew.CheckDB(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")

	router.HandleFunc("/followuser", middlew.CheckDB(middlew.ValidateJWT(routers.CreateRelations))).Methods("POST")
	router.HandleFunc("/unfollowuser", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteRelations))).Methods("DELETE")

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
