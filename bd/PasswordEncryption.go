package bd

import "golang.org/x/crypto/bcrypt"

//PasswordEncryption encrypts the user pass usin bcrypt library
func PasswordEncryption(pass string) (string, error) {
	//number of layer for encryption algo
	cost := 8
	//GeneratesFormPassword only accepts a slice of bytes []byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
