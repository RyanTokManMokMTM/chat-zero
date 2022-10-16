package model

type Room struct {
	ID   uint   `gorm:"primary_key"` //room id
	Name string `gorm:"null"`

	Users []User `gorm:"many2many:users_rooms;"`
	DefaultModel
}
