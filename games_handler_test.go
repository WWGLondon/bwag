package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func gameSetup() (
	rw *httptest.ResponseRecorder,
	r *http.Request,
	games *GamesHandler) {

	rw = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/", nil)
	games = &GamesHandler{}

	return
}

func TestWhenUserMakesGetRequestReturnAllGames(t *testing.T) {
	rw, r, games := gameSetup()

	games.ServeHTTP(rw, r)

	var g []Game
	json.NewDecoder(rw.Body).Decode(&g)

	if len(g) != 2 {
		t.Fatalf("expected two objects, got %v", len(g))
	}
}

func TestWhenUserDoesNotUseGETReturnsMethodNotAllowed(t *testing.T) {
	t.Fail()
}
