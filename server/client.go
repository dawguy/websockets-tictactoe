package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10 // Must be less than pongWait based on gorilla example
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     unsafeCheck,
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

type Response struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

type Message struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type Move struct {
	X          int `json:"x"`
	Y          int `json:"y"`
	PlayerTurn int `json:"player"`
}

// client to hub
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	// c.conn.SetReadLimit // Sets how large a message can be
	// c.conn.SetReadDeadline // Allows us to set a timeout for a connection
	// c.conn.SetPongHandler // Allows us to handle pongs in a custom way

	for {
		_, message, err := c.conn.ReadMessage()

		fmt.Printf("%v message\n", message)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}

		var res Response

		err = json.Unmarshal(message, &res)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Name: %v\n", res.Name)

		switch res.Name {
		case "place":
			var move Move

			err = json.Unmarshal(res.Data, &move)
			if err != nil {
				fmt.Println(err.Error())
			}

			// Handle the move here.
			c.hub.place <- &move
		default:
			return
		}
	}
}

// hub to client
func (c *Client) writePump() {
	// ticker := time.NewTicker(pingPeriod) // Implement write deadlines if worried about them.

	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			fmt.Printf("%v", message)
			if !ok {
				// The hub closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage) // Might need to change to json
			if err != nil {
				return
			}

			w.Write(message)
			n := len(c.send)

			for i := 0; i < n; i++ {
				w.Write(newLine)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// All collection of memory references by doing all work in new goroutines
	go client.writePump()
	go client.readPump()
}
