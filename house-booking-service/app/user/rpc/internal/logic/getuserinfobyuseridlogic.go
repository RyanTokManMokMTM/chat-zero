package logic

import (
	"context"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUserIDLogic {
	return &GetUserInfoByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByUserIDLogic) GetUserInfoByUserID(in *pb.GetUserAuthByUserIdReq) (*pb.GetUserAuthByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserAuthByUserIdResp{}, nil
}
