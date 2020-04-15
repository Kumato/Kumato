package handlers

import (
	"github.com/kumato/kumato/internal/auth"
	"net/http"
)

func getMe(w http.ResponseWriter, r *http.Request) {
	t, err := getToken(r)
	if err != nil {
		handleErr(w, http.StatusForbidden, err)
		return
	}

	handleOK(w, t.User)
}

func getToken(r *http.Request) (auth.JWTClaims, error) {
	cookie, err := r.Cookie(cookieTokenName)
	if err != nil {
		return auth.JWTClaims{}, err
	}

	token, err := auth.Parse(cookie.Value)
	if err != nil {
		return auth.JWTClaims{}, err
	}

	return token, nil
}
