package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	promLog "github.com/prometheus/common/log"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		promLog.Errorln("WS upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, inBytes, err := conn.ReadMessage()
		if err != nil {
			promLog.Errorln("WS read:", err)
			break
		}
		promLog.Warnf("WS read: %s", inBytes)

		outMsg := fmt.Sprintf("Hello world '%s'", string(inBytes))

		err = conn.WriteMessage(websocket.TextMessage, []byte(outMsg))
		if err != nil {
			promLog.Errorln("WS write:", err)
			break
		}
	}
}
