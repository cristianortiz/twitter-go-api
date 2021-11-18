package bd

import (
	"context"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateNewUser id the function to insert in DB the validated data of a new user
func CreateNewUser(u models.User) (string, bool, error) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut

	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("User")

	u.Password, _ = PasswordEncryption(u.Password)
	//insert the user data in DB collection
	result, err := col.InsertOne(ctx, u)
	//if fails, return empty string, false and the error returned by DB
	if err != nil {
		return "", false, err
	}
	//get the id of the new register from the DB collection
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//return the id as a string, a true bool and nil error
	return ObjID.String(), true, nil
}
