package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dawguy/tictactoe/game"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func start(w http.ResponseWriter, req *http.Request) {
	game.Reset()
	game.Draw()
}

func place(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var point game.Point
	err = json.Unmarshal(b, &point)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	game.PlaceValue(point.X, point.Y)
	game.Draw()

	if game.CheckWin() {
		winChar := game.GetWinner()
		fmt.Printf("%v won the game!", winChar)
	} else if game.BoardFull() {
		fmt.Printf("Tie Game\n")
	}
}

func draw(w http.ResponseWriter, req *http.Request) {
	game.Draw()
}

func main() {
	http.HandleFunc("/game/start", start)

	// Both are valid aliases for place
	http.HandleFunc("/game/move", place)
	http.HandleFunc("/game/place", place)

	http.HandleFunc("/game/draw", draw)

	http.ListenAndServe(":8090", nil)
}
