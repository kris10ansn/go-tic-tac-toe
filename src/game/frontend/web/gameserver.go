package web

import (
	"log"
	"net/http"
)

type GameServer struct {
	games            map[string]*WebGame
	gamesBroadcaster *Broadcaster
}

func CreateServer() *GameServer {
	return &GameServer{
		games:            map[string]*WebGame{},
		gamesBroadcaster: CreateBroadcaster(),
	}
}

func (server *GameServer) Run() {
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/socket/games", server.handleGamesSocket)
	http.HandleFunc("/socket/game", server.handleGameSocket)

	http.HandleFunc("/game/list", server.handleListGames)
	http.HandleFunc("/game/add", server.handleAddGame)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (server *GameServer) addGame(name string) string {
	game := CreateGame(name)

	server.games[game.Id] = game

	server.gamesBroadcaster.Broadcast(&WebsocketMessage{
		Type: MessageTypeAddGame,
		Data: game,
	})

	return game.Id
}
