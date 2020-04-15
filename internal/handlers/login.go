package handlers

import (
	"github.com/kumato/kumato/internal/auth"
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	l := loginRequest{}

	if err := parseJSON(r.Body, &l, []string{}); err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}

	token, exp, err := auth.Login(l.Username, l.Password)
	if err != nil {
		handleErr(w, http.StatusForbidden, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieTokenName,
		Value:    token,
		Expires:  time.Unix(exp, 0),
		HttpOnly: true,
	})

	handleOK(w, responseMessage{"ok", ""})
}
