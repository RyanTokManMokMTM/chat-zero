package serverWs

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type ChannelMap struct {
	sync.Mutex
	channels   map[uint]*ClientConn
	register   chan *ClientConn
	unRegister chan *ClientConn
	broadcast  chan *Message //send to all user is online - chat
}

func NewChannelMap() *ChannelMap {
	return &ChannelMap{
		channels:   make(map[uint]*ClientConn, 100),
		register:   make(chan *ClientConn),
		unRegister: make(chan *ClientConn),
		broadcast:  make(chan *Message),
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

		case message := <-ch.broadcast:
			logx.Info("send message")
			send, _ := json.Marshal(message)
			if message.ToUser > 0 {
				//Send to User
				if client, ok := ch.channels[message.UserID]; ok {
					client.send <- send
				}
			} else if message.GroupMembers != nil && message.GroupID > 0 {
				for _, id := range message.GroupMembers {
					if client, ok := ch.channels[id]; ok {
						if client.UserID == message.UserID {
							continue
						}
						client.send <- send
					}
				}
			} else {
				//send to all user
				for _, client := range ch.channels {
					client.send <- send
				}
			}
		}

	}
}
