package routers

import (
	"errors"
	"strings"

	"github.com/cristianortiz/twitter-go-api/bd"
	"github.com/cristianortiz/twitter-go-api/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//exportable vars from validated JWT claims to be used across all project packages
var Email string
var userID string

//ProcessToken validates the token and extract the values from the JWT, put the error object at the end of return parameters in func declarations
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	//slice of bytes key defined in jwt.go
	key := []byte("twitterGoAPI")
	//the struct to check the JWT must be a pointer
	claims := &models.Claim{}
	//the standard JWT encryption starts with Bearer. split it
	// splitToken := strings.Split(tk, "Bearer")
	// //check if the JWT is standard
	// if len(splitToken) != 2 {
	// 	return claims, false, string(""), errors.New("error in token structure")
	// }
	//triming the entrypted part of the token
	tk = strings.TrimSpace(tk)
	//validate token with jwt functions
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	//if token is validated, check if the user data in claims exists
	if err == nil {
		//in this case only need to know if the user exists in DB, because Email and ID are in the claims
		_, founded, _ := bd.UserExists(claims.Email)
		if founded {
			//asign the validated claims user data to exportables vars
			Email = claims.Email
			userID = claims.ID.Hex()
		}
		//this is the return when all checks are passed
		return claims, founded, userID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	//return values if err != nil
	return claims, false, string(""), err
}
