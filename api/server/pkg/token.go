package pkg

import (
	"github.com/golang-jwt/jwt"
)

func GenerateToken(accountId uint32, role string, jwtKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims(jwt.MapClaims)
	// claims["accountId"] = accountId
	// claims["role"] = role

	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
