package socket

import (
	"log"
	"time"

	"encoding/json"

	"github.com/gorilla/websocket"
)

type ClientList []*Client

type Client struct {
	conn      *websocket.Conn
	room      *Room
	ch        chan []byte
	email     string
	collegeID string
	name      string
}

const (
	pongInterval = 5 * time.Second
	pingInterval = (pongInterval * 9) / 10
)

func NewClient(conn *websocket.Conn, email string, collegeID string, name string) *Client {
	SocketLogger.Info("New Client Created: " + email)
	return &Client{conn: conn, room: nil, ch: make(chan []byte), email: email, collegeID: collegeID, name: name}
}

func (c *Client) Reader() {
	defer func() {
		c.conn.Close()
	}()

	c.conn.SetCloseHandler(func(code int, text string) error {
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
		SocketLogger.Info("Client Connection Closed")
		c.room.Close()
		return nil
	})

	err := c.conn.SetReadDeadline(time.Now().Add(pongInterval))

	if err != nil {
		log.Println(err)
		return
	}

	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			c.conn.Close()
			break
		}
		var msg Event
		err = json.Unmarshal(payload, &msg)
		if err != nil {
			log.Println("Unable to unmarshal msg:", err)
		}
		if msg.Type == initEvent {
			SocketLogger.Info("Init event received")
		} else if msg.Type == exchangeEvent {
			SocketLogger.Debug("Exchange event received")
			c.room.Send(msg.Data, c)
		}
	}
}

func (c *Client) Writer() {
	defer func() {
		c.conn.Close()
	}()

	t := time.NewTicker(pingInterval)
	defer t.Stop()
	for {
		select {
		case msg, ok := <-c.ch:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Unable to send msg:", err)
			}
			// log.Println("Message sent:", msg)
		case <-t.C:
			SocketLogger.Debug("Sending ping to " + c.name)
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				SocketLogger.Error("Unable to send ping: " + err.Error())
				return
			}

		}

	}
}

func (c *Client) pongHandler(string) error {
	SocketLogger.Debug("pong recieved from " + c.name)
	return c.conn.SetReadDeadline(time.Now().Add(pongInterval))
}
