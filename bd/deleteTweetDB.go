package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteTweetDB func delete a specific tweet in DB
func DeleteTweetDB(tweetID string, userID string) error {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("tweets")
	//convert the string id into objID for mongo
	objID, _ := primitive.ObjectIDFromHex(tweetID)
	condition := bson.M{
		"_id":    objID,
		"userID": userID,
	}
	_, err := col.DeleteOne(ctx, condition)
	return err

}
