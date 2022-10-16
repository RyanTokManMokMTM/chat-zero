package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/config"
	"time"
)

type DefaultModel struct {
	CreatedAt time.Time      `json:"-" gorm:"type:timestamp"`
	UpdatedAt time.Time      `json:"-" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp" json:"-"`
}

func NewDBEngine(c *config.Config) *gorm.DB {
	conn, err := gorm.Open(mysql.Open(c.MySQL.Datasource), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db, err := conn.DB()
	if err != nil {
		db.Close()
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}

	db.SetConnMaxIdleTime(time.Minute)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	err = conn.AutoMigrate(&User{})
	if err != nil {
		logx.Info(err)
	}

	err = conn.AutoMigrate(&Room{})
	if err != nil {
		logx.Info(err)
	}

	err = conn.AutoMigrate(&UsersRooms{})
	if err != nil {
		logx.Info(err)
	}

	err = conn.SetupJoinTable(&User{}, "Rooms", &UsersRooms{})
	if err != nil {
		logx.Info(err)
	}
	err = conn.SetupJoinTable(&Room{}, "Users", &UsersRooms{})
	if err != nil {
		logx.Info(err)
	}

	return conn
}
