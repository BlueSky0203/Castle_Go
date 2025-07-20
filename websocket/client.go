package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允許所有來源（測試階段）
		return true
	},
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v, URL: %s, Headers: %v", err, r.URL, r.Header)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	hub.Register <- client

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.Broadcast <- message
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()
	for {
		select {
		// 等待 hub 從 c.send channel 發來訊息
		case message, ok := <-c.send:
			if !ok {
				// hub 關閉了 channel，告訴 client 斷線
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}
