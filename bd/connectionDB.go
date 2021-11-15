package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongoconn exportable DBconnection object
var MongoConn = DBconnect()

//intern DB connection details
var clientOptions = options.Client().ApplyURI("mongodb+srv://cristianortiznavia:Katanakrei28@cluster0.qn9cd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//DBconnect() return a mongo client object type
func DBconnect() *mongo.Client {
	//make the connection using clientOptions data,
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		//return the client object even is empty
		return client
	}
	//check if the db is running, err= assign value to a existing value
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		//return the client object even is empty
		return client
	}
	log.Println("DB connection is running..")
	//return a valid DB connection
	return client
}

//check the DB with a ping
func CheckConnection() bool {
	//check if the db is running, in a new err variable
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true

}
