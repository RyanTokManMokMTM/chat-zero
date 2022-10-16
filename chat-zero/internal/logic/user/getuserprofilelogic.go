package user

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/ctxtool"

	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserprofileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserprofileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserprofileLogic {
	return &GetuserprofileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserprofileLogic) Getuserprofile(req *types.GetProfileReq) (resp *types.GetProfileResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCtx(l.ctx)
	logx.Info(userId)
	if userId == 0 {
		return nil, fmt.Errorf("user id not found")
	}

	u, err := l.svcCtx.DAO.UserFindOneByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &types.GetProfileResp{
		Name:  u.Name,
		Email: u.Email,
	}, nil
}
