package socket

import (
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type InterManager struct {
	queue  chan Client
	rooms  RoomList
	memory ClientList
	sync.RWMutex
}

func (m *InterManager) HandleConnections(c *gin.Context) {
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

func (m *InterManager) AddClient(client *Client) {
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

func (m *InterManager) RoomDispatcher() {
	for client := range m.queue {
		SocketLogger.Debug("Room Dispatcher dealing with client: " + client.name)
		if len(m.memory) == 0 {
			m.memory = append(m.memory, &client)
			continue
		}

		if m.memory[0].email == client.email {
			m.memory = append(m.memory, &client)
			m.memory = m.memory[1:]
			continue
		}

		room := NewRoom(m.memory[0], &client, m)
		m.addRoom(room)

		// empty memory
		m.memory = m.memory[1:]
		room.Start()
		m.addRoom(room)
	}
}

func (m *InterManager) addRoom(room *Room) {
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

func (m *InterManager) RemoveRoom(room *Room) {
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

func (m *InterManager) Close() {
	// m.Lock()
	close(m.queue)
	for room := range m.rooms {

		room.Close()
	}
	// m.Unlock()
}
