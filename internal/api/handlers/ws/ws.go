package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//ServeWs set set websocket connection
func ServeWs(w http.ResponseWriter, r *http.Request) {
	//upgrade
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Error().Msgf("websocket handshake error: %s", err.Error())
			//todo:
		}
		return
	}
	defer ws.Close()
	ws.WriteMessage(websocket.TextMessage, []byte("websocket connection set"))
	//todo:
	//add conn to clients pool
	//begin to listen messages from clients???
}
