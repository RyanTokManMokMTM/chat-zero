package dao

import (
	"context"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/model"
)

func (d *DAO) ExistInTheRoom(ctx context.Context, userID, roomID uint) error {
	ur := model.UsersRooms{
		RoomID: roomID,
		UserID: userID,
	}
	return ur.FindOne(d.engine, ctx)
}

func (d *DAO) GetRoomUsers(ctx context.Context, roomID uint) ([]uint, error) {
	ur := model.UsersRooms{
		RoomID: roomID,
	}
	return ur.GetRoomUsers(d.engine, ctx)
}
