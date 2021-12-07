package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllUsersDB get all the users in DB, but if receieve the parameter relType "R" will get
//only those users related to the user logged in the app
func GetAllUsersDB(ID string, page int64, search string, relType string) ([]*models.User, bool) {
	//define a new context for this operation and add it on top of inicial context
	//(context.Background) but this mini context will exists only for 15 secs max
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer to executes this line as the last instruction in this function
	defer cancel() //close the new mini context canceling the timeOut
	db := MongoConn.Database("twitter-go-DB")
	col := db.Collection("User")
	//the slice of users that will be returned in http response
	var results []*models.User

	//options to config the query in Find()
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	condition := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	cursor, err := col.Find(ctx, condition, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false

	}
	var founded, include bool
	//looping true the cursor with the results from DB
	for cursor.Next(ctx) {
		//every element in cursor is a User struct of models.User type
		//we need to assign the results to the same type to work on it
		var s models.User
		//decode every user
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		//if there is no error to read a a user data from cursor, time to get the user relations
		var r models.Relations
		r.UserID = ID //the global var from jwt
		//get the ID of followed user and convert to string from the mongo ID format
		r.FollowedUserID = s.ID.Hex()
		include = false
		founded, _ = GetUserRelationsDB(r)
		//for a new user that is not being followed
		if relType == "new" && !founded {
			include = true

		}
		if relType == "follow" && founded {
			include = true

		}
		//in case the ID received belongs to the logged user
		if r.FollowedUserID == ID {
			include = false
		}
		//for any user includes we don't need this data fields
		if include {
			s.Password = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			//add to results slice, the current user in cursor
			results = append(results, &s)
		}

	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)
	return results, true
}
