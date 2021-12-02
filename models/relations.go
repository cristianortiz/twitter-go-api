package models

//Relations struct for the relations of follow another user by their id
type Relations struct {
	UserID         string `bson:"userID" json:"userID"`
	FollowedUserID string `bson:"followedUserID" json:"followedUserID"`
}
