package model

import (
	"context"
	"gorm.io/gorm"
)

type Message struct {
	ID     uint `gorm:"primaryKey"`
	RoomID uint
	Sender uint
	Data   string
	DefaultModel

	RoomInfo Room `gorm:"foreignKey:RoomID;references:ID"`
	SendUser User `gorm:"foreignKey:Sender;references:ID"`
}

/*
A -> RoomID x
x ,A,"message1",time
x ,B,"message2",time
x ,C,"message3",time

sorting by time
*/

func (m *Message) TableName() string {
	return "message"
}

func (m *Message) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&m).Error
}

func (m *Message) GetRoomMessages(db *gorm.DB, ctx context.Context) ([]Message, error) {
	var record []Message
	if err := db.WithContext(ctx).Debug().Where("room_id = ?", m.RoomID).Find(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}
