package helpers

import (
	"Golang-FGA-FinalProject/database"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(database.GoDotEnvVariable("SECRET_KEY")))

	return signedToken
}

func VerifyToken(tokenStr string) (interface{}, error) {
	errResponse := errors.New("token invalid")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(database.GoDotEnvVariable("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
