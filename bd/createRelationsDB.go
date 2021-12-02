package bd

import (
	"context"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
)

//CreateRelationsDB insert in db a new models.Relations type data in DB
func CreateRelationsDB(t models.Relations) (bool, error) {
	//(context.Background) but this mini context will exists only for 15 secs max
	//define a new context for this operation and add it on top of inicial context
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("relations")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
