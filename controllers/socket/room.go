package socket

import (
	"encoding/json"
)

type Room struct {
	Name    string
	clients ClientList
	Manager Manager
}

type RoomList map[*Room]bool

func NewRoom(c1 *Client, c2 *Client, m Manager) *Room {
	return &Room{
		Name:    c1.email + c2.email,
		clients: ClientList{c1, c2},
		Manager: m,
	}
}

func (r *Room) Start() {
	for _, client := range r.clients {
		go client.Writer()
		go client.Reader()
	}
	SocketLogger.Info("Room Started with " + r.clients[0].email + " and " + r.clients[1].email)
	r.Send(json.RawMessage([]byte(`{"type": "init", "user": "`+r.clients[1].name+`","email": "`+r.clients[1].email+`"}`)), r.clients[0])
}

func (r *Room) Send(data json.RawMessage, sender *Client) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case string:
				SocketLogger.Error("Recovered from panic: " + r)
			case error:
				SocketLogger.Error("Recovered from panic: " + r.Error())
			default:
				SocketLogger.Error("Recovered from panic")
			}
		}
	}()
	for _, client := range r.clients {
		if client != sender {
			// send message data to client
			SocketLogger.Debug("Sending message " + string(data) + " to " + client.email + " from " + sender.email)
			client.ch <- data
		}
	}
}

func (r *Room) Close() {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case string:
				SocketLogger.Error("Recovered from panic: " + r)
			case error:
				SocketLogger.Error("Recovered from panic: " + r.Error())
			default:
				SocketLogger.Error("Recovered from panic")
			}
		}
	}()
	for _, client := range r.clients {
		close(client.ch)
		client.conn.Close()
	}
	r.Manager.RemoveRoom(r)

}
