package socket

import (
	"encoding/json"
)

type Room struct {
	Name    string
	clients ClientList
	Manager *Manager
}

type RoomList map[*Room]bool

func NewRoom(c1 *Client, c2 *Client, m *Manager) *Room {
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
	r.Send(Event{
		Type: openEvent,
		Data: json.RawMessage(`{"message": "Room Opened"}`),
	}, r.clients[0])
}

func (r *Room) Send(event Event, sender *Client) {
	for _, client := range r.clients {
		if client != sender {
			// send event to client
			client.ch <- event.Data
		}
	}
}

func (r *Room) Close() {
	for _, client := range r.clients {
		client.conn.Close()
		close(client.ch)
	}
	r.Manager.RemoveRoom(r)
}
