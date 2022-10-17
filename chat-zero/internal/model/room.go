package model

import (
	"context"
	"gorm.io/gorm"
)

type Room struct {
	ID       uint   `gorm:"primary_key"` //room id
	Name     string `gorm:"null"`        //room name
	Info     string `gorm:"null"`        //room info
	OwnerRef uint
	Owner    User   `gorm:"foreignKey:OwnerRef"`
	Users    []User `gorm:"many2many:users_rooms;"`
	DefaultModel
}

/*
Single Chat(Peer to Peer)

*/

func (r *Room) TableName() string {
	return "room"
}

func (r *Room) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&r).Error
}
func (r *Room) FindOne() {}
