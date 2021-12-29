package web

import (
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

type Game struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	playerX *websocket.Conn
	playerO *websocket.Conn
}

func (g *Game) AwaitMove(board *game.Board, turn game.Tic) (byte, byte) {
	panic("Not implemented")
}

func (g *Game) PresentBoard(board *game.Board) {
	panic("Not implemented")
}

func (g *Game) EndGame(board *game.Board, winner game.Tic, moves byte) {
	panic("Not implemented")
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

	g.setPlayer(tic, conn)

	if g.playerX != nil && g.playerO != nil {
		g.Start()
	}

	return nil
}

func (g *Game) setPlayer(tic game.Tic, conn *websocket.Conn) {
	var p **websocket.Conn

	if tic == game.X_TIC {
		p = &g.playerX
	} else if tic == game.O_TIC {
		p = &g.playerO
	} else {
		panic(fmt.Sprintf("Unknown tic %d", tic))
	}

	*p = conn

	conn.WriteJSON(WebsocketMessage{
		Type: "assign-tic",
		Data: game.TicToString(tic),
	})
}
