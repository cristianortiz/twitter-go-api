package models

import (
	"time"
)

//Tweet struct for insert i DB
type TweetsModel struct {
	UserID     string    `bson:"userID" json:"userID,omitempty"`
	Message    string    `bson:"message" json:"message,omitempty"`
	Created_at time.Time `bson:"created_at" json:"created_at,omitempty"`
}
