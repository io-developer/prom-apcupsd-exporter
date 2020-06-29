package server

import (
	"encoding/json"
	"fmt"
	"local/apcupsd_exporter/model"
	"net/http"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/websocket"
)

const (
	wsDefaultWriteWait      = 15 * time.Second
	wsDefaultPongWait       = 10 * time.Second
	wsDefaultPingPeriod     = 5 * time.Second
	wsDefaultMaxMessageSize = int64(64 * 1024)
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	wsClients         = map[*WsClient]bool{}
	wsUnregisterQueue = make(chan *WsClient)
)

// wsInit ..
func wsInit() {
	go wsListenUnregister()

	changeQueue := make(chan *model.Model)
	collector.GetModel().AddOnChange(changeQueue)
	go wsListenModelChange(changeQueue)

	http.HandleFunc("/ws", wsOnConnect)
}

// HandleWs ..
func wsOnConnect(w http.ResponseWriter, r *http.Request) {
	level.Debug(logger).Log("msg", "Incoming websocket connection")

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		level.Error(logger).Log("msg", "connection upgrade error", "err", err)
		return
	}

	client := NewWsClient(conn, wsUnregisterQueue)
	wsClients[client] = true

	wsSendInit(client)
}

func wsListenUnregister() {
	for {
		client, ok := <-wsUnregisterQueue
		if !ok {
			level.Warn(logger).Log("msg", "stop listening for clients unregistering")
			return
		}
		if _, exists := wsClients[client]; exists {
			delete(wsClients, client)
		} else {
			level.Warn(logger).Log("msg", "unregister faile: client not registered")
		}
	}
}

func wsSendInit(client *WsClient) {
	payload := map[string]interface{}{
		"type":        "init",
		"message":     "Init complete. Listening UPS events..",
		"model_state": collector.GetModel().State,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		level.Error(logger).Log("msg", "init payload jsonErr", "err", err)
		return
	}
	client.sendQueue <- WsMsg{
		msgType: websocket.TextMessage,
		data:    payloadJSON,
	}
}

func wsBroadcast(msgType int, msgData []byte) {
	level.Debug(logger).Log("msg", fmt.Sprintf(
		"broadcasting msg to %d connections", len(wsClients),
	))
	for client := range wsClients {
		client.sendQueue <- WsMsg{
			msgType: msgType,
			data:    msgData,
		}
	}
}

func wsListenModelChange(ch chan *model.Model) {
	for {
		if model, ok := <-ch; ok {
			wsOnModelChange(model)
		} else {
			return
		}
	}
}

func wsOnModelChange(m *model.Model) {
	level.Debug(logger).Log(
		"msg", "ws onModelChange",
		"diff", fmt.Sprintf("%#v", m.ChangedFields),
	)
	payload := map[string]interface{}{
		"type":             "change",
		"model_state_diff": m.ChangedFields,
	}
	if jsonStr, err := json.Marshal(payload); err == nil {
		wsBroadcast(websocket.TextMessage, jsonStr)
	} else {
		level.Warn(logger).Log("msg", "onModelChange json error", "err", err)
	}
}

// WsMsg ..
type WsMsg struct {
	msgType int
	data    []byte
}

// WsClient ..
type WsClient struct {
	conn           *websocket.Conn
	unregister     chan *WsClient
	sendQueue      chan WsMsg
	writeWait      time.Duration
	pongWait       time.Duration
	pingPeriod     time.Duration
	maxMessageSize int64
}

// NewWsClient ..
func NewWsClient(conn *websocket.Conn, unregister chan *WsClient) *WsClient {
	c := &WsClient{
		conn:           conn,
		unregister:     unregister,
		sendQueue:      make(chan WsMsg),
		writeWait:      wsDefaultWriteWait,
		pongWait:       wsDefaultPongWait,
		pingPeriod:     wsDefaultPingPeriod,
		maxMessageSize: wsDefaultMaxMessageSize,
	}
	go c.listenRead()
	go c.listenSend()
	return c
}

func (c *WsClient) listenRead() {
	defer func() {
		level.Debug(logger).Log("msg", "listenRead unregistering")
		c.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(c.maxMessageSize)
	c.conn.SetPongHandler(func(string) error {
		level.Debug(logger).Log("msg", "listenRead SetPongHandler")
		c.conn.SetReadDeadline(time.Now().Add(c.pongWait))
		return nil
	})
	c.conn.SetReadDeadline(time.Now().Add(c.pongWait))

	for {
		msgType, msgData, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				level.Debug(logger).Log("msg", "listenRead unexpected close error", "err", err)
			}
			return
		}
		c.onReadMsg(WsMsg{
			msgType: msgType,
			data:    msgData,
		})
	}
}

func (c *WsClient) listenSend() {
	ticker := time.NewTicker(c.pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {

		case msg, ok := <-c.sendQueue:
			c.conn.SetWriteDeadline(time.Now().Add(c.writeWait))
			if !ok {
				level.Debug(logger).Log("msg", "listenSend conn closed")
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(msg.msgType)
			if err != nil {
				level.Debug(logger).Log("msg", "listenSend nextWiter error", "err", err)
				return
			}

			w.Write(msg.data)
			if err := w.Close(); err != nil {
				level.Debug(logger).Log("msg", "listenSend writer closing error", "err", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(c.writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				level.Debug(logger).Log("msg", "listenSend ping error", "err", err)
				return
			}
		}
	}
}

func (c *WsClient) sendMsg(msg WsMsg) {
	level.Debug(logger).Log("msg", "sending msg to client")

	err := c.conn.WriteMessage(msg.msgType, msg.data)
	if err != nil {
		level.Debug(logger).Log("msg", "sendMsg error, unregistering client", "err", err)
		c.unregister <- c
	}
}

func (c *WsClient) onReadMsg(msg WsMsg) {
	level.Debug(logger).Log("msg", "onReadMsg", "type", msg.msgType, "text", string(msg.data))
}
