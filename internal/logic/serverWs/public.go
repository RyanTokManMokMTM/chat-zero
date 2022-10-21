package serverWs

import (
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func SendNotificationToUser(from, to uint, data string) error {
	logx.Info("send add friend notification...")
	message := &Message{
		Type:         SYSTEM,
		GroupID:      0,
		ToUser:       to,
		UserID:       from,
		Content:      data,
		SendTime:     time.Now().Unix(),
		GroupMembers: nil,
	}
	globalHub.broadcast <- message
	return nil
}
