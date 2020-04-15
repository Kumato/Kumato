package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var privateKey = []byte("9z60do57vk9LxpA4wV0k1dJKW35GAPRR")

const issuer = "MKKII"

func (jc *JWTClaims) token() (string, error) {
	expiresAt := time.Now().AddDate(0, 1, 0)
	jc.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiresAt.Unix(),
		Issuer:    issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jc)
	return token.SignedString(privateKey)
}

func Parse(token string) (JWTClaims, error) {
	t, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		return JWTClaims{}, err
	}

	claims, ok := t.Claims.(*JWTClaims)
	if !ok {
		return JWTClaims{}, errors.New("invalid token")
	}

	if err := claims.Valid(); err != nil {
		return JWTClaims{}, err
	}

	return *claims, err
}
