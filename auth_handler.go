package main

import "net/http"

// AuthHandler is a middleware which checks the authorisation before calling the
// next handler in the chain
type AuthHandler struct {
	Next http.Handler
}

func (a *AuthHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Nic" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	a.Next.ServeHTTP(rw, r)
}
