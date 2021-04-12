package middlewares

import (
	"net/http"
	"strings"

	"github.com/SKilliu/novels-backend/internal/errs"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

var (
	userID      = "user_id"
	accountType = "account_type"
)

func (m Middleware) ParseToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, err := GetUserIDFromJWT(c.Request(), m.auth)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, errs.UnauthorizedErr)
		}

		return next(c)
	}
}

func GetUserIDFromJWT(r *http.Request, auth string) (string, string, error) {
	return getFromJWT(r, auth, userID)
}

func GetAccountTypeFromJWT(r *http.Request, auth string) (string, string, error) {
	return getFromJWT(r, auth, accountType)
}

func getFromJWT(r *http.Request, auth string, fieldType string) (string, string, error) {
	var tokenRaw string
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		tokenRaw = bearer[7:]
	}

	token, err := jwt.Parse(tokenRaw, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(auth), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("cannot cast token.Claims to jwt.MapClaims")
	}

	var fieldRaw interface{}

	fieldRaw, ok = claims[fieldType]
	if !ok {
		return "", "", errors.New("user info is absent in the jwt")
	}

	fieldValue, ok := fieldRaw.(string)
	if !ok {
		return "", "", errors.New("failed to cast shopper_id into string")
	}

	return fieldValue, token.Raw, nil
}
