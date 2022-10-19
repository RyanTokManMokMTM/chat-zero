package serverWs

import (
	"context"
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

func NewClientConn(userID uint, conn *websocket.Conn, hub *ChannelMap, svcCtx *svc.ServiceContext) *ClientConn {
	return &ClientConn{
		hub:    hub,
		UserID: userID,
		conn:   conn,
		send:   make(chan []byte),
		svcCtx: svcCtx,
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
		req := &MessageReq{}
		err = json.Unmarshal(msg, req)
		if err != nil {
			logx.Error(err)
			//there may send back an error message
			continue
		}
		u, err := c.svcCtx.DAO.UserFindOneByID(context.TODO(), c.UserID)
		if err != nil {
			logx.Error(err)
			continue
		}
		//TODO: Get Room ID From JSON
		if err := c.svcCtx.DAO.ExistInTheRoom(context.TODO(), c.UserID, req.GroupID); err != nil {
			logx.Error(err)
			continue
		}
		//TODO: Check User
		//TODO: Store Message
		if err := c.svcCtx.DAO.InsertOneMessage(context.TODO(), req.GroupID, c.UserID, req.Message); err != nil {
			logx.Error()
			continue
		}
		//TODO: send the message to all user to all room user who is online
		allUser, err := c.svcCtx.DAO.GetRoomUsers(context.TODO(), req.GroupID)
		if err != nil {
			logx.Error(err)
			continue
		}

		message := &Message{
			Type:    MESSAGE,
			GroupID: req.GroupID,
			ToUser:  0,
			UserID:  c.UserID,
			UserDetail: SenderData{
				UserID:   u.ID,
				UserName: u.Name,
			},
			Content:      req.Message,
			SendTime:     time.Now().Unix(),
			GroupMembers: allUser,
		}

		c.hub.broadcast <- message
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
		case data, ok := <-c.send:
			/*
				TODO:
				Response:
				1. Type of data - system or message
				2. UserSent
				3. Data ： message

			*/

			//set  write deadline and send
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * WriteWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logx.Error(err)
				return
			}

			_, _ = w.Write(data)
			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(data)
			}

			if err := w.Close(); err != nil {
				logx.Error(err)
				return
			}

		case <-t.C:
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * 45))
			//send a ping message
		}
	}
}
