package logic

import (
	"context"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *pb.SignInReq) (*pb.SignInResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SignInResp{}, nil
}
