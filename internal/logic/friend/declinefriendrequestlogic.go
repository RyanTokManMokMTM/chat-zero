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

type DeclineFriendRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeclineFriendRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeclineFriendRequestLogic {
	return &DeclineFriendRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeclineFriendRequestLogic) DeclineFriendRequest(req *types.DeclineFriendNotificationReq) (resp *types.DeclineFriendNotificationResp, err error) {
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

	//TODO: Check request is exist or request state is ture
	_, err = l.svcCtx.DAO.FindOneFriendNotificationByID(l.ctx, req.RequestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("friend request not found")
		}
		return nil, err
	}

	//TODO: Set The notification state to false
	if err := l.svcCtx.DAO.DeclineFriendNotification(l.ctx, req.RequestID); err != nil {
		return nil, err
	}

	return &types.DeclineFriendNotificationResp{
		Message: fmt.Sprintf("friend request is declined"),
	}, nil
}
