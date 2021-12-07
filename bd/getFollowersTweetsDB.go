package bd

import (
	"context"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFollowersTweetsDB(ID string, page int) ([]*models.FollowersTweetsModel, bool) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("relations")

	skip := (page - 1) * 20
	//multiple query conditions make a M type slice
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userID": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userID",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"created_at": -1}})
	conditions = append(conditions, bson.M{"skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})
	//Aggregate, new Mongo framework runs the query in DB
	cursor, _ := col.Aggregate(ctx, conditions)
	var result []*models.FollowersTweetsModel
	//cursor.All read and process the cursor elements and put them in result var
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true

}
