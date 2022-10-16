package serverWs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type ChannelMap struct {
	sync.Mutex
	channels      map[uint]*ClientConn
	register      chan *ClientConn
	unRegister    chan *ClientConn
	broadcast     chan *MessageReq //send to all user is online
	systemMessage chan []byte
}

func NewChannelMap() *ChannelMap {
	return &ChannelMap{
		channels:      make(map[uint]*ClientConn, 100),
		register:      make(chan *ClientConn),
		unRegister:    make(chan *ClientConn),
		broadcast:     make(chan *MessageReq),
		systemMessage: make(chan []byte),
	}
}

func (ch *ChannelMap) Add(id uint, client *ClientConn) (*websocket.Conn, bool) {
	//add a new client to map
	//here we need to disconnect the old channel later
	conn, ok := ch.channels[id]
	ch.channels[id] = client
	if ok {
		return conn.conn, ok
	}
	return nil, ok
}

func (ch *ChannelMap) Remove(id uint) {
	//remove an existing client from map
	ch.Lock()
	defer ch.Unlock()
	delete(ch.channels, id)
}

func (ch *ChannelMap) Run() {
	//receiving sign
	for {
		select {
		case client := <-ch.register:
			conn, ok := ch.Add(client.UserID, client)

			if ok {
				conn.WriteMessage(websocket.CloseMessage, nil)
				//break the connection
				conn.Close()
			}

		case client := <-ch.unRegister:
			logx.Info("Client left!")
			ch.Remove(client.UserID)

		case msg := <-ch.broadcast:
			logx.Info("send message")
			if client, ok := ch.channels[msg.ToUser]; ok {
				logx.Infof("Sent message to user %d", msg.ToUser)
				client.conn.WriteMessage(websocket.TextMessage, []byte(msg.Message))
			}
		case msg := <-ch.systemMessage:
			info := fmt.Sprintf("[SYSTEM MESSAGE] : %v", string(msg))
			for _, client := range ch.channels {
				client.conn.WriteMessage(websocket.TextMessage, []byte(info))
			}
		}

	}
}
