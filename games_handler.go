package main

import (
	"encoding/json"
	"net/http"
)

// Game represent a game which is for sale
type Game struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Players int
	Type    string
	Price   float64
}

var games = []Game{
	Game{
		ID:   1,
		Name: "Risk",
	},
	Game{
		ID:   2,
		Name: "Dr Lucky",
	},
}

// GamesHandler does some stuff for handling games
type GamesHandler struct{}

func (g *GamesHandler) ServeHTTP(
	rw http.ResponseWriter, r *http.Request) {

	json.NewEncoder(rw).Encode(games)

	/*
	* {
	*   "Name": "Risk",
	*   ...
	* }
	 */

}
