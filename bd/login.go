package bd

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/cristianortiz/twitter-go-api/models"
)

//Login check the user  password in DB to login a valid user
func Login(email string, password string) (models.User, bool) {

	user, founded, _ := UserExists(email)
	if !founded {
		return user, false
	}
	//password in []byte type, to use with bcrypt
	passwordBytes := []byte(password)
	//password returned by DB from UserExists()
	passwordDB := []byte(user.Password)
	//bcrypt hashes the password passed to Login as a parameter
	//ah then compare with hashed pass returned by DB
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	//if pass are differnet return false, the user data does not matter
	if err != nil {
		return user, false
	}
	return user, true
}
