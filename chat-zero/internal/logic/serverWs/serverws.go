package serverWs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/ctxtool"
	"net/http"
)

type MessageReq struct {
	FromUser uint
	ToUser   uint   `json:"to_user"`
	GroupID  uint   `json:"group_id"`
	Message  string `json:"message"`
}

type MessageResp struct {
	IsSystem uint   `json:"is_system"`
	FromUser uint   `json:"to_user"`
	Message  []byte `json:"message"`
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
	client := NewClientConn(userId, conn, hub)
	hub.register <- client
	hub.systemMessage <- []byte(fmt.Sprintf("%s logged in", u.Name))

	go client.ReadLoop()
	go client.WriteLoop()

}
