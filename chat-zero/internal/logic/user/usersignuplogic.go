package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/crytox"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/ctxtool"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/jwtx"
	"time"

	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersignupLogic {
	return &UsersignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersignupLogic) Usersignup(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.UserFindOneByEmail(l.ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	encryptedPw := crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt)

	res, err := l.svcCtx.DAO.UserInsertOne(l.ctx, req.Email, encryptedPw, req.Name)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	expiredTime := now + l.svcCtx.Config.Auth.AccessExpire
	payLoad := map[string]any{
		ctxtool.CTXJWTUserID: res.ID,
	}
	token, err := jwtx.GenerateToken(expiredTime, now, l.svcCtx.Config.Auth.AccessSecret, payLoad)
	if err != nil {
		return nil, err
	}

	return &types.SignUpResp{
		Token:       token,
		ExpiredTime: uint(expiredTime),
	}, nil
}
