package svc

import (
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/config"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
