package main

import (
	"bytes"
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

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
		c.hub.broadcast <- message
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
