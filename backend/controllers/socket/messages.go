package socket

type ChatMessage struct {
	RoomID  string `json:"roomID"`
	Message string `json:"message"`
	Time    string `json:"time"`
}
