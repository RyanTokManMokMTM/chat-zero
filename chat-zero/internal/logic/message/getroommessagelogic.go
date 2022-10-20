package message

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

type GetRoomMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomMessageLogic {
	return &GetRoomMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomMessageLogic) GetRoomMessage(req *types.GetRoomMessageReq) (resp *types.GetRoomMessageResp, err error) {
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

	//TODO: Check User is joined the group
	mem, err := l.svcCtx.DAO.FindOneRoomMember(l.ctx, req.RoomID, u.ID)
	if err != nil {
		return nil, err
	}

	if mem.ID == 0 {
		return nil, fmt.Errorf("you haven't joined the group")
	}

	//TODO: Get at most 10 latest record belong to the group
	msgs, err := l.svcCtx.DAO.GetRoomMessage(l.ctx, req.RoomID)
	if err != nil {
		return nil, err
	}

	var record []types.MessageData
	for _, data := range msgs {
		record = append(record, types.MessageData{
			MessageID: data.ID,
			UserInfo: types.SenderInfo{
				UserID:   data.SendUser.ID,
				UserName: data.SendUser.Name,
			},
			Content:  data.Data,
			SendTime: data.CreatedAt.Unix(),
		})
	}

	return &types.GetRoomMessageResp{
		Messagees: record,
	}, nil
}
