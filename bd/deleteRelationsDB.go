package bd

import (
	"context"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
)

//DeleteRelationsDB drops the relations of an user following another in DB
func DeleteRelationsDB(t models.Relations) (bool, error) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("relations")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil

}
