package websocket

import (
	"encoding/json"
	"log"
)

const (
	messageTypeMove         = "move"
	messageTypeAwaitingMove = "await-move"
	messageTypePresentBoard = "present-board"
	messageTypeEndGame      = "end-game"
)

type websocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type websocketMoveMessage struct {
	Type string `json:"type"`
	Move move   `json:"data"`
}

func (p *websocketPlayer) sendMessage(messageType string, data interface{}) error {
	return p.conn.WriteJSON(websocketMessage{messageType, data})
}

func (p *websocketPlayer) readMessages() {
	for {
		_, messageString, err := p.conn.ReadMessage()

		if err != nil {
			log.Println("Error: ", err)
			break
		}

		var message *websocketMessage
		err = json.Unmarshal(messageString, &message)

		if err != nil {
			log.Println(err)
			continue
		}

		switch message.Type {
		case messageTypeMove:
			var moveMessage *websocketMoveMessage
			err = json.Unmarshal(messageString, &moveMessage)

			if err != nil {
				log.Println(err)
				continue
			}

			p.moves <- moveMessage.Move
		default:
			log.Printf("Message with unknown type \"%s\" received", message.Type)
		}
	}
}
