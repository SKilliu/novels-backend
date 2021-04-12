package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type tokenField string

var (
	userID      tokenField = "user_id"
	accountType tokenField = "account_type"
)

func GenerateJWT(uid, accType, verifyKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		string(userID):      uid,
		string(accountType): accType,
	})

	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

