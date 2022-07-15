package logic

import (
	"context"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByAuthKeyLogic {
	return &GetUserInfoByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByAuthKeyLogic) GetUserInfoByAuthKey(in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserAuthByAuthKeyResp{}, nil
}
