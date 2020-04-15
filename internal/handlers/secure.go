package handlers

import (
	"github.com/kumato/kumato/internal/auth"
	"net/http"
)

func secure(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieTokenName)
		if err != nil {
			// handleErr(w, http.StatusForbidden, err)
			http.Redirect(w, r, r.URL.Scheme+"://"+r.URL.Host+"/login", http.StatusForbidden)
			return
		}

		_, err = auth.Parse(cookie.Value)
		if err != nil {
			// handleErr(w, http.StatusForbidden, err)
			http.Redirect(w, r, r.URL.Scheme+"://"+r.URL.Host+"/login", http.StatusForbidden)
			return
		}

		http.HandlerFunc(next).ServeHTTP(w, r)
	})
}

func secureInternal(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("INTERNAL_TOKEN") != auth.InternalToken() {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		http.HandlerFunc(next).ServeHTTP(w, r)
	})
}
