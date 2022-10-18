package dao

import (
	"context"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/model"
)

func (d *DAO) InsertOneRoom(ctx context.Context, name, info string, userID uint) (*model.Room, error) {
	r := &model.Room{
		Name:     name,
		Info:     info,
		OwnerRef: userID,
	}

	if err := r.InsertOne(d.engine, ctx); err != nil {
		return nil, err
	}

	return r, nil
}

func (d *DAO) RemoveOneRoom(ctx context.Context, roomID uint) error {
	r := &model.Room{
		ID: roomID,
	}

	return r.RemoveOne(d.engine, ctx)
}

func (d *DAO) FindOneOwnerRoom(ctx context.Context, roomID, userID uint) (*model.Room, error) {
	r := &model.Room{
		ID:       roomID,
		OwnerRef: userID,
	}

	if err := r.FindOne(d.engine, ctx); err != nil {
		return nil, err
	}

	return r, nil
}

func (d *DAO) FindOneByRoomID(ctx context.Context, roomID uint) (*model.Room, error) {
	r := &model.Room{
		ID: roomID,
	}

	if err := r.FindOne(d.engine, ctx); err != nil {
		return nil, err
	}

	return r, nil
}

func (d *DAO) JoinOneRoom(ctx context.Context, roomID uint, u *model.User) error {
	r := model.Room{
		ID: roomID,
	}
	return r.InsertOneUser(d.engine, ctx, u)
}

func (d *DAO) LeaveOneRoom(ctx context.Context, roomID uint, u *model.User) error {
	r := model.Room{
		ID: roomID,
	}
	return r.RemoveOneUser(d.engine, ctx, u)
}

func (d *DAO) FindRoomMembers(ctx context.Context, roomID uint) ([]*model.User, error) {
	r := model.Room{
		ID: roomID,
	}
	return r.FindRoomMembers(d.engine, ctx)
}
