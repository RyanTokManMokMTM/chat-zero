package serverWs

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"sync"
	"time"
)

type ClientConn struct {
	sync.Mutex
	hub *ChannelMap

	UserID uint
	conn   *websocket.Conn
	svcCtx *svc.ServiceContext
	send   chan []byte
}

type Message struct {
	From uint
	To   uint
	Type uint   // 0() or 1
	Data string //actual data from user
}

func NewClientConn(userID uint, conn *websocket.Conn, hub *ChannelMap) *ClientConn {
	return &ClientConn{
		hub:    hub,
		UserID: userID,
		conn:   conn,
		send:   make(chan []byte),
	}
}

func (c *ClientConn) ReadLoop() {
	defer func() {
		c.hub.unRegister <- c // remove client from map
		c.conn.Close()        //close connection
	}()

	c.conn.SetReadDeadline(time.Now().Add(time.Second * ReadWait))
	c.conn.SetReadLimit(ReadLimit)
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(time.Second * ReadWait))
		return nil
	})

	for {
		//get data from connection
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			logx.Error(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Error(err)
			}
			break
		}
		//TODO: record : userID ,ToUerID, message, sendTime
		req := &MessageReq{}
		err = json.Unmarshal(msg, req)
		req.FromUser = c.UserID
		if err != nil {
			logx.Error(err)
		}

		if req.ToUser > 0 {
			if _, ok := c.hub.channels[req.ToUser]; !ok {
				logx.Error("User %v is offline\n", req.ToUser)
				continue
			}
		}

		c.hub.broadcast <- req
	}
}

func (c *ClientConn) WriteLoop() {
	t := time.NewTicker(time.Second * (WriteWait / 2))
	defer func() {
		c.hub.unRegister <- c // remove client from map
		c.conn.Close()        //close connection
	}()

	for {
		select {
		case msg, ok := <-c.send:
			//set  write deadline and send
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * WriteWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)     //send the first message
			n := len(c.send) //any other message else?

			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-t.C:
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * 45))
			//send a ping message
		}
	}
}
