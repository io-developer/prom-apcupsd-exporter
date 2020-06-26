package server

import (
	"fmt"
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
	wsOnMsgPushCh     = make(chan clientMsg)
	wsOnModelChangeCh = make(chan *model.Model)
)

// RegisterWsEndpoints ..
func RegisterWsEndpoints(c *metric.Collector) {
	collector = c

	collector.GetModel().AddOnChange(wsOnModelChangeCh)

	go listenMsgPush(wsOnMsgPushCh)
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

	wsOnMsgPushCh <- clientMsg{
		mtype: websocket.TextMessage,
		data:  []byte("Init complete. Listening UPS events.."),
		conn:  conn,
	}
}

// Broadcast ..
func Broadcast(text string) {
	broadcast(websocket.TextMessage, []byte(text))
}

func broadcast(msgType int, msgData []byte) {
	for conn := range wsConnections {
		wsOnMsgPushCh <- clientMsg{
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

func listenMsgPush(ch chan clientMsg) {
	for {
		if msg, ok := <-ch; ok {
			onMsgPush(msg)
		} else {
			return
		}
	}
}

func onMsgPush(msg clientMsg) {
	err := msg.conn.WriteMessage(msg.mtype, msg.data)
	if err != nil {
		delete(wsConnections, msg.conn)
		promLog.Errorln("WS write:", err)
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

	str := "Changes:\n"
	for field, diff := range m.ChangedFields {
		str += fmt.Sprintf("  '%s'\n    OLD: %#v\n    NEW: %#v\n", field, diff[0], diff[1])
	}

	Broadcast(str)
}
