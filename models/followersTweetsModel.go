package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FollowersTweetsModel struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	FollowerUserID   string             `bson:"followerUserID" json:"followerUserID,omitempty"`
	FollowerRelation string             `bson:"followerRelation" json:"followerRelation,omitempty"`
	Tweet            struct {
		Message    string    `bson:"message" json:"message,omitempty"`
		Created_at time.Time `bson:"created_at" json:"created_at,omitempty"`
		ID         string    `bson:"_id" json:"_id,omitempty"`
	}
}
