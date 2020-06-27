package server

import (
	"encoding/json"
	"fmt"
	"local/apcupsd_exporter/metric"
	"local/apcupsd_exporter/model"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/websocket"
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
	level.Debug(Logger).Log("msg", "Incoming websocket connection")

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		level.Error(Logger).Log("msg", "connection upgrade error", "err", err)
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
		level.Error(Logger).Log("msg", "init payload jsonErr", "err", err)
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
	level.Debug(Logger).Log("msg", fmt.Sprintf(
		"broadcasting msg to %d connections", len(wsConnections),
	))
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
	level.Debug(Logger).Log("msg", "send msg to client")

	err := msg.conn.WriteMessage(msg.mtype, msg.data)
	if err != nil {
		level.Error(Logger).Log("msg", "pushMsg error, removing bad connection from list", "err", err)

		delete(wsConnections, msg.conn)
		defer msg.conn.Close()
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
	level.Debug(Logger).Log(
		"msg", "ws onModelChange",
		"diff", fmt.Sprintf("%#v", m.ChangedFields),
	)

	payload := map[string]interface{}{
		"type":             "change",
		"model_state_diff": m.ChangedFields,
	}
	if jsonStr, err := json.Marshal(payload); err == nil {
		broadcast(websocket.TextMessage, jsonStr)
	} else {
		level.Warn(Logger).Log("msg", "onModelChange json error", "err", err)
	}
}
