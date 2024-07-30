package chat_websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}


func Upgrader(w http.ResponseWriter ,r *http.Request)(*websocket.Conn,error){
	upgrade.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	connection, err:= upgrade.Upgrade(w, r, nil)
		if err != nil{
			log.Println("Websocket connectionn Error :--",err)
			return nil ,err
		}
		return connection ,nil
	
}