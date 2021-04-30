package jwt

import (
	"fmt"
	"time"

	models "github.com/Thiti-Dev/tawb-service-v1/models/jwt"
	"github.com/dgrijalva/jwt-go"
)

// GetSignedTokenFromData -> is a (fn) that will sign a token with the given data
func GetSignedTokenFromData(data models.RequiredDataToClaims) string{
	claims := models.CustomClaims{
		Username: data.Username,
		ID: data.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "aittty.io",
		},
	}
	// Create an unsigned token from the claims above
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token ->  preferably at least 256 bits in length (in-production xD)
	signedToken, err := token.SignedString([]byte("TemporaryJustForNow"))
	if err != nil {
		fmt.Println(err)
	}
	return signedToken
}

func DecodeJwt(tokenString string) (bool,interface{}){
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    	return []byte("TemporaryJustForNow"), nil
	})
	if err != nil{
		return false, nil
	}
	return true,claims
}