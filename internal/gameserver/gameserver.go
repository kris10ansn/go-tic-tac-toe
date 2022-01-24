package gameserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/kris10ansn/go-tic-tac-toe/internal/players/socket"
	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type GameServer struct {
	games map[string]*game.Game
}

var upgrader websocket.Upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

func (server *GameServer) joinGame(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	gameId := parameters["id"]

	conn, err := upgrader.Upgrade(rw, r, nil)

	if g, exists := server.games[gameId]; exists {
		if err != nil {
			log.Println(err)
		}

		player := socket.CreateWebsocketPlayer(conn)
		tic, err := g.AddPlayer(player)

		if err != nil {
			log.Println(err)
		}

		if tic != game.EMPTY_TIC {
			player.AssignTic(tic)
		}

		if !g.Playing && g.GetPlayer(game.X_TIC) != nil && g.GetPlayer(game.O_TIC) != nil {
			g.Play()
		}
	} else {
		message := []byte(fmt.Sprintf("__Game with id '%s' does not exist", gameId))
		conn.WriteMessage(websocket.CloseMessage, message)
		conn.Close()
	}
}

func Host() {
	server := &GameServer{
		games: make(map[string]*game.Game),
	}
	server.games["sex"] = &game.Game{}

	r := mux.NewRouter()

	r.HandleFunc("/ws/game/{id}", server.joinGame)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
