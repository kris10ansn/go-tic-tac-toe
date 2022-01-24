package socket

import (
	"github.com/gorilla/websocket"
	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type move struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type WebsocketPlayer struct {
	conn  *websocket.Conn
	moves chan move
	tic   game.Tic
}

func CreateWebsocketPlayer(conn *websocket.Conn) *WebsocketPlayer {
	player := &WebsocketPlayer{
		conn:  conn,
		moves: make(chan move),
	}

	go player.readMessages()
	return player
}

func (p *WebsocketPlayer) AwaitMove(board *game.Board, turn game.Tic) (int, int) {
	p.sendMessage(messageTypeAwaitingMove, p.tic)
	move := <-p.moves

	return move.X, move.Y
}
func (p *WebsocketPlayer) Present(board *game.Board) {
	p.sendMessage(messageTypePresentBoard, board)
}
func (p *WebsocketPlayer) EndGame(winner game.Tic, moves byte) {
	p.sendMessage(messageTypeEndGame, struct {
		Winner game.Tic `json:"winner"`
		Moves  byte     `json:"moves"`
	}{winner, moves})
}

func (p *WebsocketPlayer) AssignTic(tic game.Tic) {
	p.tic = tic
}
