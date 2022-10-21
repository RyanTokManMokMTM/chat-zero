package dao

import (
	"context"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/model"
)

func (d *DAO) UserInsertOne(ctx context.Context, email, password, name string) (*model.User, error) {
	u := &model.User{
		Email:    email,
		Password: password,
		Name:     name,
	}

	if err := u.InsertOne(d.engine, ctx); err != nil {
		return nil, err
	}
	return u, nil
}

func (d *DAO) UserFindOneByID(ctx context.Context, userID uint) (*model.User, error) {
	u := &model.User{
		ID: userID,
	}
	if err := u.FindOneByID(d.engine, ctx); err != nil {
		return nil, err
	}

	return u, nil
}

func (d *DAO) UserFindOneByEmail(ctx context.Context, email string) (*model.User, error) {
	u := &model.User{
		Email: email,
	}
	if err := u.FindOneByEmail(d.engine, ctx); err != nil {
		return nil, err
	}

	return u, nil
}
