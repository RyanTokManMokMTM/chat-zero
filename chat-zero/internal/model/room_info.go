package model

type UsersRooms struct {
	RoomID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
}

func (ur *UsersRooms) TableName() string {
	return "users_rooms"
}
