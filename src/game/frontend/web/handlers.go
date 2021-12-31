package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (server *GameServer) handleGamesSocket(rw http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(rw, r, nil, 1024, 1024)

	if err != nil {
		log.Println("Error:", err)
	}

	LogWebsocketConnection(conn, r)

	server.gamesBroadcaster.Join(conn)
	defer server.gamesBroadcaster.Leave(conn)

	for _, message, err := conn.ReadMessage(); true; {
		if err != nil {
			log.Println("Error:", err)
			return
		}

		log.Println(message)
	}
}

func (server *GameServer) handleGameSocket(rw http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(rw, r, nil, 1024, 1024)

	if err != nil {
		log.Println("Error:", err)
	}

	LogWebsocketConnection(conn, r)

	var joinMessage struct {
		Type string `json:"type"`
		Data struct {
			GameId string `json:"gameId"`
		} `json:"data"`
	}

	if err1 := conn.ReadJSON(&joinMessage); err1 != nil {
		log.Println(err1)
	}

	if game, exists := server.games[joinMessage.Data.GameId]; exists {
		joinError := game.Join(conn)

		if joinError != nil {
			conn.WriteJSON(WebsocketMessage{
				Type: MessageTypeError,
				Data: joinError.Error(),
			})
		}
	} else {
		conn.WriteJSON(WebsocketMessage{
			Type: MessageTypeError,
			Data: "game does not exist",
		})
	}
}

func (server *GameServer) handleListGames(rw http.ResponseWriter, r *http.Request) {
	gamesList := []*Game{}

	for _, game := range server.games {
		gamesList = append(gamesList, game)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(gamesList)
}

func (server *GameServer) handleAddGame(rw http.ResponseWriter, r *http.Request) {
	var game struct {
		Name string `json:"name"`
	}
	jsonError := json.NewDecoder(r.Body).Decode(&game)

	if jsonError != nil {
		http.Error(rw, jsonError.Error(), http.StatusBadRequest)
		return
	}

	id := server.addGame(game.Name)

	json.NewEncoder(rw).Encode(struct {
		Id string `json:"id"`
	}{id})

	rw.WriteHeader(http.StatusCreated)
}

func LogWebsocketConnection(conn *websocket.Conn, r *http.Request) {
	log.Printf("Websocket connection on %s [%s => %s]",
		r.URL.Path, conn.RemoteAddr(), conn.LocalAddr())
}
