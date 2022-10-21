package serverWs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/ctxtool"
	"net/http"
	"time"
)

var globalHub *ChannelMap

const (
	SYSTEM = iota
	MESSAGE
)

type SenderData struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

type MessageReq struct {
	GroupID uint   `json:"group_id"`
	Message string `json:"message"`
}

type Message struct {
	Type         uint       `json:"message_type"`
	GroupID      uint       `json:"group_id"` //for chat
	GroupMembers []uint     `json:"-"`
	ToUser       uint       `json:"to_user"` //for notification
	UserID       uint       `json:"user_id"`
	UserDetail   SenderData `json:"sender_info"`
	Content      string     `json:"content"`
	SendTime     int64      `json:"send_time"`
}

const (
	ReadWait  = 60
	WriteWait = 60
	ReadLimit = 1024
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewServerWS(svcCtx *svc.ServiceContext) func(w http.ResponseWriter, r *http.Request) {
	if globalHub == nil {
		logx.Info("initialing hub...")
		globalHub = NewChannelMap()
		go globalHub.Run()
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ServerWS(svcCtx, globalHub, w, r)
	}
}

func ServerWS(ctx *svc.ServiceContext, hub *ChannelMap, w http.ResponseWriter, r *http.Request) {

	userId := ctxtool.GetUserIDFromCtx(r.Context())
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := ctx.DAO.UserFindOneByID(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logx.Infof("Client %d(name:%s) is connected via websocket", userId, u.Name)
	client := NewClientConn(userId, conn, hub, ctx)
	hub.register <- client
	hub.broadcast <- &Message{
		Type:         SYSTEM,
		GroupID:      0,
		ToUser:       0,
		UserID:       u.ID,
		Content:      fmt.Sprintf("[SYSTEM] %s is now online.", u.Name),
		SendTime:     time.Now().Unix(),
		GroupMembers: nil,
	}

	go client.ReadLoop()
	go client.WriteLoop()

}
