package svc

import (
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/model/user"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     user.UserModel
	UserAuthModel user.UserAuthModel
	RedisCache    cache.CacheConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.Datasource)
	return &ServiceContext{
		Config:        c,
		UserModel:     user.NewUserModel(conn, c.RedisCache),
		UserAuthModel: user.NewUserAuthModel(conn, c.RedisCache),
	}
}
