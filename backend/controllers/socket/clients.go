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
}

const (
	pongInterval = 10 * time.Second
	pingInterval = (pongInterval * 9) / 10
)

func NewClient(conn *websocket.Conn, manager *Manager, email string, collegeID string) *Client {
	return &Client{conn: conn, room: nil, ch: make(chan []byte), email: email, collegeID: collegeID}
}

func (c *Client) Reader() {
	defer func() {
		c.room.Close()
	}()

	err := c.conn.SetReadDeadline(time.Now().Add(pongInterval))

	if err != nil {
		log.Println(err)
		return
	}

	c.conn.SetPongHandler(c.pongHandler)

	for {

		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			c.room.Close()
			break
		}
		var msg Event
		err = json.Unmarshal(payload, &msg)
		if err != nil {
			log.Println("Unable to unmarshal msg:", err)
		}

		if msg.Type == chatMessageEvent {
			var chatMsg ChatMessageEventData
			err = json.Unmarshal(msg.Data, &chatMsg)
			if err != nil {
				log.Println("Unable to unmarshal data:", err)
			}

			payload, err = json.Marshal(chatMsg)
			if err != nil {
				log.Println("Unable to marshal msg:", err)
			}

		}

		var event Event
		event.Data = payload
		event.Type = msg.Type

		c.room.Send(event, c)

	}
}

func (c *Client) Writer() {
	defer func() {
		c.room.Close()
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
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				SocketLogger.Error("Unable to send ping: " + err.Error())
				return
			}

		}

	}

}

func (c *Client) pongHandler(string) error {
	return c.conn.SetReadDeadline(time.Now().Add(pongInterval))
}
