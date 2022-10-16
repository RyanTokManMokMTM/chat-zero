package dao

import "gorm.io/gorm"

type DAO struct {
	engine *gorm.DB
}

func NewDAO(db *gorm.DB) *DAO {
	return &DAO{
		engine: db,
	}
}
