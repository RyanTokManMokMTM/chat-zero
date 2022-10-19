package friend

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

type GetFriendRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendRequestLogic {
	return &GetFriendRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendRequestLogic) GetFriendRequest(req *types.GetFriendRequestReq) (resp *types.GetFriendRequestResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCtx(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	_, err = l.svcCtx.DAO.UserFindOneByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	list, err := l.svcCtx.DAO.GetFriendRequest(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	var request []types.FriendRequest
	for _, req := range list {
		request = append(request, types.FriendRequest{
			RequestID: req.ID,
			Sender:    req.Sender,
		})
	}
	return &types.GetFriendRequestResp{
		Requests: request,
	}, nil
}
