package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SearchProfile search and find in DB the data of a logged user
func SearchProfile(ID string) (models.User, error) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	//set db and collection from DB
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("User")

	var profile models.User
	//the Id is in a string obj, we need to converto to objID to search in DB
	objID, _ := primitive.ObjectIDFromHex(ID)

	searchCondition := bson.M{
		"_id": objID,
	}
	//find the user in User DB collection
	err := col.FindOne(ctx, searchCondition).Decode(&profile)
	//avoid the password at return the data, remember the omit empty in DB collection config
	profile.Password = ""
	if err != nil {
		fmt.Println("Register not found " + err.Error())
		return profile, err
	}
	return profile, err
}
