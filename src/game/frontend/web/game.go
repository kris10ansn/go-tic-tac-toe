package web

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

type Player struct {
	conn *websocket.Conn
}

func CreatePlayer(conn *websocket.Conn) *Player {
	return &Player{
		conn:  conn,
	}
}

type Game struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	playerX *Player
	playerO *Player
}

func CreateGame(name string) *Game {
	return &Game{
		Name: name,
		Id:   uuid.NewString(),
	}
}

func (g *Game) AwaitMove(board *game.Board, turn game.Tic) (byte, byte) {
	panic("Not implemented")
}

func (g *Game) PresentBoard(board *game.Board) {
	g.writePlayers(&WebsocketMessage{
		Type: MessageTypePresentBoard,
		Data: board,
	})
}

func (g *Game) EndGame(board *game.Board, winner game.Tic, moves byte) {
	g.writePlayers(&WebsocketMessage{
		Type: MessageTypeEndGame,
		Data: struct {
			board  *game.Board
			winner game.Tic
			moves  byte
		}{board, winner, moves},
	})
}

func (g *Game) Start() {
	log.Printf("[%s] starting game...", g.Id)
}

func (g *Game) Join(conn *websocket.Conn) error {
	var tic game.Tic

	if g.playerX == nil {
		tic = game.X_TIC
	} else if g.playerO == nil {
		tic = game.O_TIC
	} else {
		return errors.New("game is full")
	}

	*(g.playerProperty(tic)) = CreatePlayer(conn)

	conn.WriteJSON(WebsocketMessage{
		Type: MessageTypeAssignTic,
		Data: game.TicToString(tic),
	})

	if g.playerX != nil && g.playerO != nil {
		g.Start()
	}

	return nil
}

func (g *Game) writePlayers(message *WebsocketMessage) {
	g.playerX.conn.WriteJSON(message)
	g.playerO.conn.WriteJSON(message)
}

func (g *Game) playerProperty(tic game.Tic) **Player {
	if tic == game.X_TIC {
		return &g.playerX
	} else if tic == game.O_TIC {
		return &g.playerO
	} else {
		panic(fmt.Sprintf("Unknown tic %d", tic))
	}
}
