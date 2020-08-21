package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Reset()
	Draw()
}

func place(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var point Point
	err = json.Unmarshal(b, &point)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	PlaceValue(point.X, point.Y)
	Draw()

	if CheckWin() {
		winChar := GetWinner()
		fmt.Printf("%v won the game!", winChar)
	} else if BoardFull() {
		fmt.Printf("Tie Game\n")
	}
}

func draw(w http.ResponseWriter, req *http.Request) {
	Draw()
}

func main() {
	http.HandleFunc("/game/start", start)

	// Both are valid aliases for place
	http.HandleFunc("/game/move", place)
	http.HandleFunc("/game/place", place)

	http.HandleFunc("/game/draw", draw)

	http.ListenAndServe(":8090", nil)
}
