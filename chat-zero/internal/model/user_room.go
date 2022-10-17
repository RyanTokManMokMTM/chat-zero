package model

import (
	"context"
	"gorm.io/gorm"
)

type UsersRooms struct {
	RoomID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
}

//There may a lot of user inside the same room
/*
For example
RoomID : x1
A :x1
B :x1
C :x1

that means A,B and C in the same chat group


*/

func (ur *UsersRooms) TableName() string {
	return "users_rooms"
}

func (ur *UsersRooms) GetRoomUsers(db *gorm.DB, ctx context.Context) ([]uint, error) {
	var allUser []uint
	if err := db.WithContext(ctx).Debug().Model(ur).Select("user_id").Where("room_id = ?", ur.RoomID).Find(&allUser).Error; err != nil {
		return nil, err
	}

	return allUser, nil
}
func (ur *UsersRooms) FindOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().First(&ur).Error
}
