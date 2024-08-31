package socket

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     checkOrigin,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	clients map[string]ClientList
	queue   chan Client
	rooms   RoomList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: map[string]ClientList{},
		queue:   make(chan Client),
		rooms:   RoomList{},
	}
}

func (m *Manager) HandleConnections(c *gin.Context) {
	SocketLogger.Info("Handling Connection")
	email := c.GetString("email")
	collegeFormat := strings.Split(email, "@")[1]

	SocketLogger.Info("Email, College in Header: " + email + " " + collegeFormat)

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		SocketLogger.Error("Problem in upgrade: " + err.Error())
	}

	name := c.GetString("name")

	client := NewClient(ws, m, email, collegeFormat, name)

	SocketLogger.Info("New Client Created")

	m.AddClient(client)
}

func (m *Manager) AddClient(client *Client) {
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

func (m *Manager) RoomDispatcher() {
	for client := range m.queue {
		_, exist := m.clients[client.collegeID]
		if !exist {
			m.clients[client.collegeID] = make([]*Client, 0)
		}
		m.clients[client.collegeID] = append(m.clients[client.collegeID], &client)
		SocketLogger.Info("Client Added to Queue")
		if len(m.clients[client.collegeID]) == 2 {
			room := NewRoom(m.clients[client.collegeID][0], m.clients[client.collegeID][1], m)
			for _, client := range m.clients[client.collegeID] {
				client.room = room
			}
			SocketLogger.Info("Room Created")
			m.addRoom(room)
			room.Start()
			delete(m.clients, client.collegeID)
		}
	}
}

func (m *Manager) addRoom(room *Room) {
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

func (m *Manager) RemoveRoom(room *Room) {
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

func checkOrigin(r *http.Request) bool {
	return true
}

func (m *Manager) Close() {
	// m.Lock()
	close(m.queue)
	for room := range m.rooms {

		room.Close()
	}
	// m.Unlock()
}
