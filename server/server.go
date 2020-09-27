package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dawguy/tictactoe/game"
)

func unsafeCheck(r *http.Request) bool {
	return true
}

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

func getBoard(w http.ResponseWriter, _ *http.Request) {
	var a = game.GetBoard()
	var board = strings.Join(a[:], ",")
	dat := []byte(board)

	w.Write(dat)
}

func serverGame() {

}

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	http.HandleFunc("/game/move", place)
	http.HandleFunc("/game/place", place)
	http.HandleFunc("/game/start", start)
	http.HandleFunc("/game/draw", draw)
	http.HandleFunc("/game/getBoard", getBoard)

	http.ListenAndServe(":8090", nil)
}
