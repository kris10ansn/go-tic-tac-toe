package web

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const BUF_SIZE = 1024

type Game struct {
	Name string
}

type GameServer struct {
	games []Game
}

func (server *GameServer) Run() {
	http.HandleFunc("/socket/games", func(rw http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(rw, r, nil, BUF_SIZE, BUF_SIZE)

		if err != nil {
			log.Println(err)
		}

		log.Printf(
			"Websocket connection on %s [%s => %s]\n",
			r.URL.Path,
			conn.RemoteAddr(),
			conn.LocalAddr(),
		)
	})

	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}

func (server *GameServer) AddGame(name string) {
	server.games = append(server.games, Game{Name: name})
}

func CreateServer() GameServer {
	return GameServer{
		games: []Game{},
	}
}
