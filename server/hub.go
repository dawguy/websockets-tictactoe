package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dawguy/tictactoe/game"
)

// Hub maintains the set o active clients. Based on gorilla example
type Hub struct {
	clients    map[*Client]bool
	place      chan *Move
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		place:      make(chan *Move),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case move := <-h.place:
			fmt.Println("PLACE CALLED")
			// Update state of game. Then broadcast new state to all clients
			game.PlaceValue(move.X, move.Y)
			a := game.GetBoard()

			var board = strings.Join(a[:], ",")
			fmt.Println(board)

			var res = Message{
				Name: "place",
				Data: board,
			}

			dat, err := json.Marshal(res)
			fmt.Printf("%v\n", string(res.Data))

			if err != nil {
				fmt.Println("HMMM")
				fmt.Println(err.Error())
				return
			}

			fmt.Println(string(dat))
			fmt.Println(res)

			for client := range h.clients {
				select {
				case client.send <- dat:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
