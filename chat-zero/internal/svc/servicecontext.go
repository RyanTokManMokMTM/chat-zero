package svc

import (
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/config"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/dao"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/model"
)

type ServiceContext struct {
	Config config.Config
	DAO    *dao.DAO
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := model.NewDBEngine(&c)
	return &ServiceContext{
		Config: c,
		DAO:    dao.NewDAO(engine),
	}
}
