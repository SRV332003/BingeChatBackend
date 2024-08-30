package socket

import "encoding/json"

type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

const (
	// Event types
	cursorMoveEvent     = "cursorMove"
	chatMessageEvent    = "chatMessage"
	openEvent           = "openEvent"
	commentMessageEvent = "commentMessage"
)

type ChatMessageEventData struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}