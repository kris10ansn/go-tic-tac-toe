package web

import "github.com/gorilla/websocket"

type Broadcaster struct {
	clients map[*websocket.Conn]bool
}

func CreateBroadcaster() *Broadcaster {
	return &Broadcaster{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *Broadcaster) Join(client *websocket.Conn) {
	s.clients[client] = true
}

func (s *Broadcaster) Leave(client *websocket.Conn) {
	delete(s.clients, client)
}

func (s *Broadcaster) Broadcast(message *WebsocketMessage) {
	for client := range s.clients {
		client.WriteJSON(message)
	}
}
