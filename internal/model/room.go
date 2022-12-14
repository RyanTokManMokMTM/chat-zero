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

func (r *Room) RemoveOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Delete(&r).Error
}

func (r *Room) FindOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().First(&r).Error
}

func (r *Room) InsertOneUser(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Debug().Model(&r).Association("Users").Append(user)
}

func (r *Room) RemoveOneUser(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Debug().Model(&r).Association("Users").Delete(user)
}

func (r *Room) FindRoomMembers(db *gorm.DB, ctx context.Context) ([]*User, error) {
	var members []*User
	err := db.WithContext(ctx).Debug().Model(&r).Association("Users").Find(&members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *Room) FindOneRoomMember(db *gorm.DB, ctx context.Context, userID uint) (*User, error) {
	var members *User
	err := db.WithContext(ctx).Debug().Model(&r).Where("user_id = ?", userID).Association("Users").Find(&members)
	if err != nil {
		return nil, err
	}
	return members, nil
}
