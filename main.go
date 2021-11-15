package main

import (
	"log"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/handlers"
)

func main() {

	if !bd.CheckConnection() {
		log.Fatal("DB is not running")
		return
	}
	handlers.ServerHandler()

}
