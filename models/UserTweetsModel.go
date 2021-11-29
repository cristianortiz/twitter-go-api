package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserTweetsModel structs represents the data structures for list all the tweets of and specific userID
type UserTweetsModel struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID     string             `bson:"userID" json:"userID,omitempty"`
	Message    string             `bson:"message" json:"message,omitempty"`
	Created_at time.Time          `bson:"created_at" json:"created_at,omitempty"`
}
