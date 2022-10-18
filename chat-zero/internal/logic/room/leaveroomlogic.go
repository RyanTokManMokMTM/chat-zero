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

type LeaveRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveRoomLogic {
	return &LeaveRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveRoomLogic) LeaveRoom(req *types.LeaveRoomReq) (resp *types.LeaveRoomResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCtx(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	//TODO: Check User is exist
	u, err := l.svcCtx.DAO.UserFindOneByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	//find the room
	//TODO: Check the room is exist
	_, err = l.svcCtx.DAO.FindOneByRoomID(l.ctx, req.RoomID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("room not exist")
		}
		return nil, err
	}

	//user already joined?
	//TODO: Check user is joined
	err = l.svcCtx.DAO.ExistInTheRoom(l.ctx, userId, req.RoomID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user is not in the room")
		}
		return nil, err
	}

	if err := l.svcCtx.DAO.LeaveOneRoom(l.ctx, req.RoomID, u); err != nil {
		return nil, err
	}

	return
}
