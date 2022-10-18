package room

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

type DeleteRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoomLogic {
	return &DeleteRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoomLogic) DeleteRoom(req *types.DeleteRoomReq) (resp *types.DeleteRoomResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCtx(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	u, err := l.svcCtx.DAO.UserFindOneByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	r, err := l.svcCtx.DAO.FindOneOwnerRoom(l.ctx, req.ID, u.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("you're not the room owner")
		}

		return nil, err
	}

	if err := l.svcCtx.DAO.RemoveOneRoom(l.ctx, r.ID); err != nil {
		return nil, err
	}

	return
}
