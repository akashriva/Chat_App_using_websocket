package main

import (
    "log"
    "net/http"

    chat_websocket "github.com/akashriva/chat_application/websocket"
)

func serveWs(pool *chat_websocket.Pool, w http.ResponseWriter, r *http.Request) {
    connection, err := chat_websocket.Upgrader(w, r)
    if err != nil {
        log.Println(err)
        return
    }
    client := &chat_websocket.Client{
        Conn: connection,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}

func setupRoutes() {
    pool := chat_websocket.NewPool()
    go pool.Start()
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {
    setupRoutes()
    log.Println("Server started on :9000")
    http.ListenAndServe(":9000", nil)
}
