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

func (d *DAO) AcceptFriendNotification(ctx context.Context, requestID uint) error {
	fr := &model.FriendNotification{
		ID: requestID,
	}

	return fr.Accept(d.engine, ctx)
}
func (d *DAO) CancelFriendNotification() {}

func (d *DAO) DeclineFriendNotification() {}

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

func (d *DAO) UpdateNotificationState(ctx context.Context, requestID uint) error {
	f := &model.Friend{
		ID: requestID,
	}

	return f.UpdateState(d.engine, ctx)
}
