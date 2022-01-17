package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type move struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type websocketPlayer struct {
	conn  *websocket.Conn
	moves chan move
}

func ConnectPlayer(conn *websocket.Conn) *websocketPlayer {
	player := &websocketPlayer{
		conn: conn,
	}

	go player.readMessages()
	return player
}

func (p *websocketPlayer) AwaitMove(board *game.Board, turn game.Tic) (int, int) {
	p.sendMessage(messageTypeAwaitingMove, nil)
	move := <-p.moves

	return move.X, move.Y
}
func (p *websocketPlayer) Present(board *game.Board) {
	p.sendMessage(messageTypePresentBoard, board)
}
func (p *websocketPlayer) EndGame(winner game.Tic, moves byte) {
	p.sendMessage(messageTypeEndGame, struct {
		Winner game.Tic `json:"winner"`
		Moves  byte     `json:"moves"`
	}{winner, moves})
}
