package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct mongo documents definition
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	LastName string             `bson:"lastname" json:"lastname,omitempty"`
	Birthday time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Location string             `bson:"location" json:"location,omitempty"`
	WebSite  string             `bson:"website" json:"website,omitempty"`
}
