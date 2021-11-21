package jwt

import (
	"time"

	"github.com/cristianortiz/twitter-go-api/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneratesJWT(t models.User) (string, error) {
	//the jwt token is an slice of bytes
	key := []byte("twitterGoAPI")
	//claims (privileges) section of the token to add in paylod
	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.Name,
		"lastname": t.LastName,
		"birthday": t.Birthday,
		"location": t.Location,
		"website":  t.WebSite,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	//header part of token, encrypton algo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//sign the token with key slice of bytes
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
