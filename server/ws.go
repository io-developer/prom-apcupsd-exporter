package server

import (
	"encoding/json"
	"local/apcupsd_exporter/metric"
	"local/apcupsd_exporter/model"
	"net/http"

	"github.com/gorilla/websocket"
	promLog "github.com/prometheus/common/log"
)

var (
	wsUpgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	wsConnections     = map[*websocket.Conn]bool{}
	wsPushMsgCh       = make(chan clientMsg)
	wsOnModelChangeCh = make(chan *model.Model)
)

// RegisterWsEndpoints ..
func RegisterWsEndpoints(c *metric.Collector) {
	collector = c

	collector.GetModel().AddOnChange(wsOnModelChangeCh)

	go listenPushMsg(wsPushMsgCh)
	go listenModelChange(wsOnModelChangeCh)

	http.HandleFunc("/ws", handleWs)
}

// HandleWs ..
func handleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		promLog.Errorln("WS upgrade:", err)
		return
	}

	wsConnections[conn] = true

	payload := map[string]interface{}{
		"type":        "init",
		"message":     "Init complete. Listening UPS events..",
		"model_state": collector.GetModel().State,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		promLog.Errorln("WS handleModelChange jsonErr", err)
		return
	}
	wsPushMsgCh <- clientMsg{
		mtype: websocket.TextMessage,
		data:  payloadJSON,
		conn:  conn,
	}
}

// Broadcast ..
func Broadcast(text string) {
	broadcast(websocket.TextMessage, []byte(text))
}

func broadcast(msgType int, msgData []byte) {
	for conn := range wsConnections {
		wsPushMsgCh <- clientMsg{
			mtype: msgType,
			data:  msgData,
			conn:  conn,
		}
	}
}

type clientMsg struct {
	mtype int
	data  []byte
	conn  *websocket.Conn
}

func listenPushMsg(ch chan clientMsg) {
	for {
		if msg, ok := <-ch; ok {
			pushMsg(msg)
		} else {
			return
		}
	}
}

func pushMsg(msg clientMsg) {
	err := msg.conn.WriteMessage(msg.mtype, msg.data)
	if err != nil {
		delete(wsConnections, msg.conn)
		promLog.Errorln("WS write:", err)
		promLog.Errorln("  removing bad connection from list")
	}
}

func listenModelChange(ch chan *model.Model) {
	for {
		if model, ok := <-ch; ok {
			onModelChange(model)
		} else {
			return
		}
	}
}

func onModelChange(m *model.Model) {
	promLog.Warnln("WS handleModelChange", m.ChangedFields)

	payload := map[string]interface{}{
		"type":             "change",
		"model_state_diff": m.ChangedFields,
	}
	if jsonStr, err := json.Marshal(payload); err == nil {
		broadcast(websocket.TextMessage, jsonStr)
	} else {
		promLog.Warnln("WS handleModelChange jsonErr", err)
	}
}
