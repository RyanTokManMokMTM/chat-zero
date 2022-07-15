package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/pb"
	"github.com/ryantokmanmokmtm/house-booking-service/common/errx"
	"github.com/ryantokmanmokmtm/house-booking-service/common/jwtx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	// todo: add your logic here and delete this line
	//Generate JWT
	nowTime := time.Now().Unix()
	tokenExp := l.svcCtx.Config.JWTAuth.AccessExpire
	tokenKey := l.svcCtx.Config.JWTAuth.AccessSecret
	token, err := jwtx.TokenGenerate(in.UserID, nowTime, tokenExp, tokenKey)
	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.TOKEN_GENERATE_ERROR), fmt.Sprintf("GenerateToken Error user: %v ,err: %v", in.UserID, err.Error()))
	}
	return &pb.GenerateTokenResp{
		AccessToken:  token,
		AccessExpire: nowTime + tokenExp,
		RefreshAfter: nowTime + tokenExp/2, //before half expired time, need to be refresh
	}, nil
}
