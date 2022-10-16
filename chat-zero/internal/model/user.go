package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	DefaultModel
	ID       uint   `gorm:"id" json:"id"`
	Email    string `gorm:"email;index:idx_email" json:"email"`
	Name     string `gorm:"name" json:"name"`
	Password string `gorm:"password" json:"password"`

	Rooms []Room `gorm:"many2many:users_rooms;"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&u).Error
}

func (u *User) FindOneByID(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().First(&u).Error
}

func (u *User) FindOneByEmail(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Where("email = ?", u.Email).First(&u).Error
}
