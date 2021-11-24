package bd

import (
	"context"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertNewTweet(t models.Tweets) (string, bool, error) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("tweets")

	record := bson.M{
		"userID":     t.UserID,
		"message":    t.Message,
		"created_at": t.Created_at,
	}
	result, err := col.InsertOne(ctx, record)
	if err != nil {
		return "", false, err
	}
	//get the id of the new tweet in objectID format
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
