package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/model/user"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/usercenter"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupLogic) Signup(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserCenterRpc.SignUp(l.ctx, &usercenter.SignUpReq{
		Email:    req.Email,
		Password: req.Password,
		AuthKey:  req.Email,
		AuthType: user.AuthTypeSystem,
	})

	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	var signUpResp types.SignUpResp
	_ = copier.Copy(&signUpResp, res)
	return &signUpResp, nil
}
