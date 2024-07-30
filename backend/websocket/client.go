package chat_websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		messageType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		m := Message{
			Type: messageType,
			Body: string(msg),
		}
		c.Pool.Broadcast <- m
		fmt.Println("Message Received: ", msg)
	}
}
