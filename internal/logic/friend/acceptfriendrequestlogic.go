package friend

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/logic/serverWs"
	"gtihub.com/ryantokmanmokmtm/chat-zero/util/ctxtool"

	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcceptFriendRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAcceptFriendRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptFriendRequestLogic {
	return &AcceptFriendRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AcceptFriendRequestLogic) AcceptFriendRequest(req *types.AcceptFriendNotificationReq) (resp *types.AcceptFriendNotificationResp, err error) {
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

	//TODO: Check request is exist or request state is ture
	notification, err := l.svcCtx.DAO.FindOneFriendNotificationByID(l.ctx, req.RequestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("friend request not found")
		}
		return nil, err
	}

	//Update Friend Request
	err = l.svcCtx.DAO.AcceptFriendNotification(l.ctx, notification)
	if err != nil {
		return nil, err
	}

	go func() {
		//Send the notification via websocket
		_ = serverWs.SendNotificationToUser(notification.Receiver, notification.Sender, fmt.Sprintf("[SYSTEM MESSAGE] %s accecpted your friend request.", u.Name))
	}()

	return &types.AcceptFriendNotificationResp{
		Message: fmt.Sprintf("friend request accepted"),
	}, nil
}
