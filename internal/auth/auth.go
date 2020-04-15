package auth

import (
	"crypto/sha512"
	"fmt"
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/logger"
	"os"
)

func Login(username, password string) (string, int64, error) {
	if m := os.Getenv("KUMATO_DEV_MODE"); m == "1" || m == "on" {
		jc := JWTClaims{
			User: db.User{
				Qid:   "n000000",
				Name:  "test user",
				Email: "test.user@test.dev",
			},
		}

		token, err := jc.token()

		return token, jc.ExpiresAt, err
	}

	user, err := authWithLDAP(username, password)
	if err != nil {
		logger.Fatal(err)
	}

	jc := JWTClaims{
		User: user,
	}

	token, err := jc.token()

	return token, jc.ExpiresAt, err
}

func InternalToken() string {
	h := sha512.New()
	h.Write(append([]byte(issuer), privateKey...))

	return fmt.Sprintf("%x", h.Sum(nil))
}
