package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAndSign(data int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": data,
		"iss":     os.Getenv("JWT_ISSUER"),
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
		// and other claims, alternatively, you may want to explore on how to create custome claims
	})

	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func ParseAndVerify(signed string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	}, jwt.WithIssuer(os.Getenv("JWT_ISSUER")),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims")
	}
}
