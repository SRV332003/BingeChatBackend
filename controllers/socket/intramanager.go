package socket

import (
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type IntraManager struct {
	clients map[string]ClientList
	queue   chan Client
	rooms   RoomList
	sync.RWMutex
}

func (m *IntraManager) HandleConnections(c *gin.Context) {
	SocketLogger.Debug("Handling Connection")
	email := c.GetString("email")
	collegeFormat := strings.Split(email, "@")[1]

	SocketLogger.Debug("Email, College in Header: " + email + " " + collegeFormat)

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		SocketLogger.Error("Problem in upgrade: " + err.Error())
	}

	name := c.GetString("name")

	client := NewClient(ws, email, collegeFormat, name)

	SocketLogger.Debug("New Client Created")

	m.AddClient(client)
}

func (m *IntraManager) AddClient(client *Client) {
	m.Lock()
	defer m.Unlock()
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
	m.queue <- *client
}

func (m *IntraManager) RoomDispatcher() {
	for client := range m.queue {
		SocketLogger.Debug("Room Dispatcher dealing with client: " + client.name)
		_, exist := m.clients[client.collegeID]
		if !exist {
			m.clients[client.collegeID] = make([]*Client, 0)
		}
		m.clients[client.collegeID] = append(m.clients[client.collegeID], &client)
		SocketLogger.Debug("Client Added to Queue")
		if len(m.clients[client.collegeID]) == 2 {
			if m.clients[client.collegeID][0].email == m.clients[client.collegeID][1].email {
				SocketLogger.Info("Same Email")
				m.clients[client.collegeID] = m.clients[client.collegeID][1:]
				continue
			}

			room := NewRoom(m.clients[client.collegeID][0], m.clients[client.collegeID][1], m)
			for _, client := range m.clients[client.collegeID] {
				client.room = room
			}
			SocketLogger.Debug("Room Created")
			m.addRoom(room)
			room.Start()
			delete(m.clients, client.collegeID)
		}
	}
}

func (m *IntraManager) addRoom(room *Room) {
	m.Lock()
	defer m.Unlock()
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
	m.rooms[room] = true
}

func (m *IntraManager) RemoveRoom(room *Room) {
	m.Lock()
	defer m.Unlock()
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
	delete(m.rooms, room)
}

func (m *IntraManager) Close() {
	// m.Lock()
	close(m.queue)
	for room := range m.rooms {

		room.Close()
	}
	// m.Unlock()
}
