package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var nextCalled = false

func setup() (
	rw *httptest.ResponseRecorder,
	r *http.Request,
	authz *AuthHandler) {

	nextCalled = false
	rw = httptest.NewRecorder()
	r = httptest.NewRequest("", "/", nil)

	authz = &AuthHandler{}
	authz.Next = http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			nextCalled = true
		},
	)

	return
}

func TestCallNextWhenAuthorized(t *testing.T) {
	rw, r, authz := setup()
	r.Header.Set("Authorization", "Nic")

	authz.ServeHTTP(rw, r)

	if nextCalled == false {
		t.Fatal("next should have been called")
	}

}

func TestDoesNotCallNextWhenNotAuthorized(t *testing.T) {
	rw, r, authz := setup()

	authz.ServeHTTP(rw, r)

	if nextCalled == true {
		t.Fatal("next should not have been called")
	}

}

func TestReturnsStatusUnauthorizedWhenNotAuthorized(t *testing.T) {
	rw, r, authz := setup()

	authz.ServeHTTP(rw, r)

	if rw.Code != http.StatusUnauthorized {
		t.Fatal("should have returned status unauthorized")
	}
}
