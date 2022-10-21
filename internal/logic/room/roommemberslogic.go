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

type RoomMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoomMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomMembersLogic {
	return &RoomMembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoomMembersLogic) RoomMembers(req *types.GetRoomMembersReq) (resp *types.GetRoomMembersResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCtx(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	//TODO: Check User is exist
	_, err = l.svcCtx.DAO.UserFindOneByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	_, err = l.svcCtx.DAO.FindOneByRoomID(l.ctx, req.RoomID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("room not exist")
		}
		return nil, err
	}

	members, err := l.svcCtx.DAO.FindRoomMembers(l.ctx, req.RoomID)
	if err != nil {
		return nil, err
	}

	var membersList []types.RoomMemberInfo
	for _, v := range members {
		membersList = append(membersList, types.RoomMemberInfo{
			UserID:   v.ID,
			UserName: v.Name,
		})
	}

	return &types.GetRoomMembersResp{
		Members: membersList,
	}, nil
}
