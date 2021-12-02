package bd

import (
	"context"
	"log"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTweetsDB, query the Tweets DB documents and returns a slice with pagination of all the tweets requested
func GetUserTweetsDB(ID string, page int64) ([]*models.UserTweetsModel, bool) {
	//(context.Background) but this mini context will exists only for 15 secs max
	//define a new context for this operation and add it on top of inicial context
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("tweets")

	var results []*models.UserTweetsModel

	condition := bson.M{
		"userID": ID,
	}
	//define the option mode of Mongo as Find() for pagination
	opt := options.Find()
	//now detail that option mode
	opt.SetLimit(20) //how many results per page
	//order by "created_at" column and in descending order, most recent first
	opt.SetSort(bson.D{{Key: "created_at", Value: -1}})
	//the first time dont skip any result, then every skip step will be on 20 results
	//per page according to the pagination number defined before
	opt.SetSkip((page - 1) * 20)
	//cursor is pointer with te DB results, can be looped and arrenged as it need it
	cursor, err := col.Find(ctx, condition, opt)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	//TODO creates an empty context to looping and work with the cursor
	for cursor.Next(context.TODO()) {
		//to keep every tweet in the cursor before using the slice defined in "results"
		var record models.UserTweetsModel
		//cursor is already a JSON because it came from the DB, so it can decoded it
		err := cursor.Decode(&record)
		if err != nil {
			return results, false
		}
		//append to result slice every referenced and decoded record from DB
		results = append(results, &record)
	}

	return results, true
}
