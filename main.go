package main

import (
	"net/http"
)

func main() {
	gamez := &GamesHandler{}
	authz := &AuthHandler{}
	authz.Next = gamez

	http.Handle("/", authz)
	http.ListenAndServe(":8085", nil)
}
