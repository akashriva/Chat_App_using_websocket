package chat_websocket

import "fmt"

type Pool struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    Broadcast  chan Message
}

func NewPool() *Pool {
    return &Pool{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan Message),
    }
}

func (pool *Pool) Start() {
    for {
        select {
        case client := <-pool.Register:
            pool.Clients[client] = true
            fmt.Println("Total connections in pool: ", len(pool.Clients))
            for client := range pool.Clients {
                client.Conn.WriteJSON(Message{Type: 1, Body: "New User joined"})
            }

        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            fmt.Println("Total connections in pool: ", len(pool.Clients))
            for client := range pool.Clients {
                client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected"})
            }

        case message := <-pool.Broadcast:
            fmt.Println("Sending message to clients!")
            for client := range pool.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                }
            }
        }
    }
}
