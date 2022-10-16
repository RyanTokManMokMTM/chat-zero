package ws

import "github.com/gorilla/websocket"

type Client struct {
	Id   int
	conn *websocket.Conn
	send chan []byte
}
