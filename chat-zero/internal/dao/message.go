package dao

import (
	"context"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/model"
)

func (d *DAO) InsertOneMessage(ctx context.Context, roomID, userId uint, message string) error {
	msg := &model.Message{
		RoomID: roomID,
		Sender: userId,
		Data:   message,
	}

	return msg.InsertOne(d.engine, ctx)
}

func (d *DAO) GetRoomMessage(ctx context.Context, roomID uint) ([]*model.Message, error) {
	msg := &model.Message{
		RoomID: roomID,
	}

	return msg.GetRoomMessages(d.engine, ctx)
}
