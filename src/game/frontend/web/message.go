package web

type WebsocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
