package socket

import "encoding/json"

type Event struct {
	Type string          `json:"identifier"`
	Data json.RawMessage `json:"data"`
}

const (
	// Event types
	exchangeEvent = "exchange"
	initEvent     = "initEvent"
)

func (e Event) ToJSON() []byte {
	data, _ := json.Marshal(e)
	return data
}
