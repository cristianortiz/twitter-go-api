package models

//Tweet captures the message form the body response
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
