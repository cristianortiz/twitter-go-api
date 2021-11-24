package bd

import (
	"context"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateUser update User collection in DB data for a specific ID
func UpdateUserDB(u models.User, ID string) (bool, error) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("User")
	//record is a map with a string key but their values are interface type
	//this is one way to make the data structure to update a collection in DB
	record := make(map[string]interface{})
	if len(u.Name) > 0 {
		record["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		record["lastname"] = u.LastName
	}
	if len(u.Email) > 0 {
		record["email"] = u.Email
	}

	record["birthday"] = u.Birthday

	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		record["website"] = u.WebSite
	}
	// part of the update query in mongoDB
	updateString := bson.M{

		"$set": record,
	}
	//convert the ID in a string to an objectID type to use in mongoDB
	objID, _ := primitive.ObjectIDFromHex(ID)
	//filter the user ID equal to the ID stored in DB to apply the update
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	//now put it all together, context, the filter and the update query
	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
