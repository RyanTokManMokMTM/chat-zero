package user

import (
	"context"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXAuthReq) (resp *types.WXNAutoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
