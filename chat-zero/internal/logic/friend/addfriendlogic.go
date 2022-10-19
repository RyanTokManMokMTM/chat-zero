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

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendReq) (resp *types.AddFriendResp, err error) {
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

	//TODO: Is Friend already?
	isFriend, err := l.svcCtx.DAO.IsFriend(l.ctx, userId, req.UserID)
	if err != nil {
		return nil, err
	}

	if isFriend {
		return &types.AddFriendResp{
			Message: fmt.Sprintf("both of your already in friendship."),
		}, nil
	}
	//TODO:Add to the notification if it hasn't sent a request.
	_, err = l.svcCtx.DAO.FindOneFriendNotification(l.ctx, userId, req.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := l.svcCtx.DAO.InsertOneFriendNotification(l.ctx, userId, req.UserID); err != nil {
			return nil, err
		}
		go func() {
			//Send the notification via websocket
			_ = serverWs.SendNotificationToUser(u.ID, req.UserID, fmt.Sprintf("[SYSTEM MESSAGE] %s sent you a friend request", u.Name))
		}()

		return &types.AddFriendResp{
			Message: fmt.Sprintf("friend request sent"),
		}, nil
	}

	return nil, fmt.Errorf("friend request had been sent")
}
