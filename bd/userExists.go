package bd

import (
	"context"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserExits is function to check if an emal o a new user is already on use by another one
//it receive an string as parmeter an return a document type from DB, a bool and a string with the id returned by DB
func UserExists(email string) (models.User, bool, string) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("User")
	//bson covert to JSON the key value pair, condition to find a specific row in DB
	condition := bson.M{"email": email}
	//models.User var to keep the returned value of bd if the email already exists
	var result models.User
	//find if the email exists in DB, decode the response to JSON and pass by reference to result
	err := col.FindOne(ctx, condition).Decode(&result)
	//convert the objID type value to string hexadecimal
	ID := result.ID.Hex()
	if err != nil {
		//if the email exists return the model.User document, FALSE and the id returned by DB
		return result, false, ID
	}
	//if the email exists return the model.User document, true and the id returned by DB
	return result, true, ID
}
