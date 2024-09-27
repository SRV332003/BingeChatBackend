package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Manager interface {
	Close()
	RemoveRoom(room *Room)
	AddClient(client *Client)
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     checkOrigin,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func checkOrigin(r *http.Request) bool {
	return true
}
