package dao

import (
	"context"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/model"
)

func (d *DAO) InsertOneFriendNotification(ctx context.Context, sender, receiver uint) error {
	fr := &model.FriendNotification{
		Sender:   sender,
		Receiver: receiver,
		State:    true,
	}

	return fr.InsertOne(d.engine, ctx)
}
func (d *DAO) FindOneFriendNotification(ctx context.Context, sender, receiver uint) (*model.FriendNotification, error) {
	fr := &model.FriendNotification{
		Sender:   sender,
		Receiver: receiver,
		State:    true,
	}
	err := fr.FineOne(d.engine, ctx)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func (d *DAO) FindOneFriendNotificationByID(ctx context.Context, requestID uint) (*model.FriendNotification, error) {
	fr := &model.FriendNotification{
		ID:    requestID,
		State: true,
	}

	err := fr.FineOne(d.engine, ctx)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func (d *DAO) AcceptFriendNotification(ctx context.Context, fr *model.FriendNotification) error {
	return fr.Accept(d.engine, ctx)
}
func (d *DAO) CancelFriendNotification(ctx context.Context, requestID uint) error {
	f := &model.FriendNotification{
		ID: requestID,
	}

	return f.Cancel(d.engine, ctx)
}

func (d *DAO) DeclineFriendNotification(ctx context.Context, requestID uint) error {
	f := &model.FriendNotification{
		ID: requestID,
	}

	return f.Decline(d.engine, ctx)
}

func (d *DAO) HasFriendShip(ctx context.Context, userID, friendID uint) (*model.User, error) {
	f := &model.Friend{
		UserID: userID,
	}

	return f.FindOneFriend(d.engine, ctx, friendID)
}

func (d *DAO) InsertOneFriendInstance(ctx context.Context, userID uint) error {
	f := &model.Friend{
		UserID: userID,
	}

	return f.InsertOne(d.engine, ctx)
}

func (d *DAO) RemoveFriend(ctx context.Context, userID, friendID uint) error {
	f := &model.Friend{}

	return f.RemoveOne(d.engine, ctx, userID, friendID)
}

func (d *DAO) GetFriendRequest(ctx context.Context, userID uint) ([]*model.FriendNotification, error) {
	notification := &model.FriendNotification{
		Receiver: userID,
		State:    true,
	}

	return notification.GetNotifications(d.engine, ctx)
}

func (d *DAO) IsFriend(ctx context.Context, userID, friendID uint) (bool, error) {
	f := &model.Friend{
		UserID: userID,
	}

	return f.IsFriend(d.engine, ctx, friendID)
}
