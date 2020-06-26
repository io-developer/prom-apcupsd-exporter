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
	initialized = false
	upgrader    = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	activeConns     = map[*websocket.Conn]bool{}
	onMsgPushCh     = make(chan clientMsg)
	onModelChangeCh = make(chan *model.Model)
)

func init() {
	if !initialized {
		initialized = true

		metric.CurrentModel.AddOnChange(onModelChangeCh)

		go listenMsgPush()
		go listenModelChange()
	}
}

// HandleWs ..
func HandleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		promLog.Errorln("WS upgrade:", err)
		return
	}

	activeConns[conn] = true

	onMsgPushCh <- clientMsg{
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
	for conn := range activeConns {
		onMsgPushCh <- clientMsg{
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

func listenMsgPush() {
	for {
		if msg, ok := <-onMsgPushCh; ok {
			onMsgPush(msg)
		} else {
			return
		}
	}
}

func onMsgPush(msg clientMsg) {
	err := msg.conn.WriteMessage(msg.mtype, msg.data)
	if err != nil {
		delete(activeConns, msg.conn)
		promLog.Errorln("WS write:", err)
	}
}

func listenModelChange() {
	for {
		if model, ok := <-onModelChangeCh; ok {
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
