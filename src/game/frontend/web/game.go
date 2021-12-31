package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

type GameMove struct {
	X byte `json:"x"`
	Y byte `json:"y"`
}

type Player struct {
	conn *websocket.Conn

	moves chan GameMove
}

func CreatePlayer(conn *websocket.Conn) *Player {
	return &Player{
		conn:  conn,
		moves: make(chan GameMove),
	}
}

type WebsocketMoveMessage struct {
	Type string   `json:"type"`
	Data GameMove `json:"data"`
}

func (p *Player) ReadMessages() {
	for {
		_, msg, err := p.conn.ReadMessage()

		if err != nil {
			log.Println("Error:", err)
			return
		}

		var message WebsocketMessage
		json.Unmarshal(msg, &message)

		switch message.Type {
		case MessageTypeClientMove:
			{
				var moveMessage WebsocketMoveMessage
				json.Unmarshal(msg, &moveMessage)

				if err != nil {
					log.Println("Error:", err)
					break
				}

				p.moves <- moveMessage.Data
			}
		default:
			{
				log.Printf("Message with unknown type \"%s\" received", message.Type)
			}
		}
	}
}

type WebGame struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	playerX *Player
	playerO *Player
}

func CreateGame(name string) *WebGame {
	return &WebGame{
		Name: name,
		Id:   uuid.NewString(),
	}
}

func (g *WebGame) AwaitMove(board *game.Board, turn game.Tic) (byte, byte) {
	player := *g.playerProperty(turn)
	player.conn.WriteJSON(WebsocketMessage{MessageTypeAwaitingMove, ""})

	move := <-player.moves
	return move.X, move.Y
}

func (g *WebGame) PresentBoard(board *game.Board) {
	g.writePlayers(&WebsocketMessage{
		Type: MessageTypePresentBoard,
		Data: board,
	})
}

func (g *WebGame) EndGame(board *game.Board, winner game.Tic, moves byte) {
	g.writePlayers(&WebsocketMessage{
		Type: MessageTypeEndGame,
		Data: struct {
			Board  *game.Board `json:"board"`
			Winner game.Tic    `json:"winner"`
			Moves  byte        `json:"moves"`
		}{board, winner, moves},
	})
}

func (g *WebGame) Start() {
	log.Printf("[%s] starting game...", g.Id)
	game.PlayGame(g)
}

func (g *WebGame) Join(conn *websocket.Conn) error {
	var tic game.Tic

	if g.playerX == nil {
		tic = game.X_TIC
	} else if g.playerO == nil {
		tic = game.O_TIC
	} else {
		return errors.New("game is full")
	}

	player := CreatePlayer(conn)
	*(g.playerProperty(tic)) = player

	go player.ReadMessages()

	conn.WriteJSON(WebsocketMessage{
		Type: MessageTypeAssignTic,
		Data: tic,
	})

	if g.playerX != nil && g.playerO != nil {
		g.Start()
	}

	return nil
}

func (g *WebGame) writePlayers(message *WebsocketMessage) {
	g.playerX.conn.WriteJSON(message)
	g.playerO.conn.WriteJSON(message)
}

func (g *WebGame) playerProperty(tic game.Tic) **Player {
	if tic == game.X_TIC {
		return &g.playerX
	} else if tic == game.O_TIC {
		return &g.playerO
	} else {
		panic(fmt.Sprintf("Unknown tic %d", tic))
	}
}
