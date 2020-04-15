package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kumato/kumato/internal/db"
)

type JWTClaims struct {
	db.User
	jwt.StandardClaims
}
