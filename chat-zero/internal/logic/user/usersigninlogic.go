package user

import (
	"context"
	"fmt"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/crytox"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/ctxtool"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/jwtx"
	"time"

	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersigninLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersigninLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersigninLogic {
	return &UsersigninLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersigninLogic) Usersignin(req *types.SignInReq) (resp *types.SignInResp, err error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.DAO.UserFindOneByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if u.Password != crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt) {
		return nil, fmt.Errorf("email or password incorrect")
	}

	now := time.Now().Unix()
	expiredTime := now + l.svcCtx.Config.Auth.AccessExpire
	payLoad := map[string]any{
		ctxtool.CTXJWTUserID: u.ID,
	}
	token, err := jwtx.GenerateToken(expiredTime, now, l.svcCtx.Config.Auth.AccessSecret, payLoad)
	if err != nil {
		return nil, err
	}
	return &types.SignInResp{
		Token:       token,
		ExpiredTime: uint(expiredTime),
	}, nil
	return
}
