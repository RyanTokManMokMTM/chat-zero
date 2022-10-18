package model

import (
	"context"
	"gorm.io/gorm"
)

type FriendNotification struct {
	ID       uint `gorm:"primary_key"`
	Sender   uint
	Receiver uint
	State    bool //0: accepted/declined/canceled 1:waiting/sent
	DefaultModel

	SenderInfo   User `gorm:"foreignKey:Sender;references:ID"`
	ReceiverInfo User `gorm:"foreignKey:Receiver;references:ID"`
}

func (f *FriendNotification) TableName() string {
	return "friend_notification"
}

func (f *FriendNotification) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&f).Error
}

func (f *FriendNotification) FineOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().First(&f).Error
}

func (f *FriendNotification) Accept(db *gorm.DB, ctx context.Context) error {
	//calling Friend model

	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		/*
			TODO: Update Notification State
			TODO: Add Friend to Friendship of both of them
		*/

		if err := tx.WithContext(ctx).Debug().Model(&f).Update("state", false).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Debug().Model(&Friend{
			UserID: f.Sender,
		}).Association("Friend").Append(&User{ID: f.Receiver}); err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Debug().Model(&Friend{
			UserID: f.Receiver,
		}).Association("Friend").Append(&User{ID: f.Sender}); err != nil {
			return err
		}

		return nil
	})
}
func (f *FriendNotification) Cancel()  {}
func (f *FriendNotification) Decline() {}
