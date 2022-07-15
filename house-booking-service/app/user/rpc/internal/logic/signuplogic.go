package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/model/user"
	"github.com/ryantokmanmokmtm/house-booking-service/common/encryptx"
	"github.com/ryantokmanmokmtm/house-booking-service/common/errx"
	"github.com/ryantokmanmokmtm/house-booking-service/common/str_generate"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(in *pb.SignUpReq) (*pb.SignUpResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err == sqlx.ErrNotFound {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("User SignIn  - Email:%s,err: %v", in.Email, err.Error()))
	}

	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.USER_HAS_BEEN_REGISTERED), fmt.Sprintf("User SignIn - Email already exists ,Email: %v,err: %v", in.Email, err))
	}

	//Using transaction?
	//Using transaction mode to ensure all transaction is completed,otherwise roll back to previous state
	var userID int64
	if err := l.svcCtx.UserModel.Transaction(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//register service
		newUser := &user.User{
			Email: in.Email,
		}

		//check username is empty?
		if len(in.Name) == 0 {
			//using a tool to generate a random name
			newUser.Name = str_generate.StrGenerator(8, str_generate.RAND_KIND_ALL)
		}
		//encrypted password -> using bcrypt tool
		if len(in.Password) > 0 {
			newUser.Password = encryptx.EncryptPassword([]byte(in.Password), []byte(l.svcCtx.Config.Salt))
		}
		//insert to model
		sqlRes, err := l.svcCtx.UserModel.Insert(ctx, newUser)
		if err != nil {
			//%+v output struct including field and value
			return errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("User SignIn - user db Insert  err:%v,user: %+v", err.Error(), newUser))
		}
		lastID, err := sqlRes.LastInsertId()
		if err != nil {
			return errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("User SignIn - user db InsertResult.LastInsertId err: %v, user:%+v", err.Error(), newUser))
		}
		userID = lastID

		//insert to user_auth
		newUserAuth := &user.UserAuth{
			UserId:   lastID,
			AuthType: in.AuthType,
			AuthKey:  in.AuthKey,
		}

		_, err = l.svcCtx.UserAuthModel.Insert(ctx, newUserAuth)
		if err != nil {
			return errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("User SignIn - user_auth db Insert err: %v ,user_auth: %+v", err.Error(), newUserAuth))
		}

		return nil
	}); err != nil {
		return nil, err //this error is our custom error
	}

	//all data is inserted successfully
	//generate the jwt token and return
	//it just simply generate the service ,other than calling an RPC service
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserID: userID,
	})
	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.TOKEN_GENERATE_ERROR), fmt.Sprintf("User SignIn - GenerateToken For UserID: %v, err:%v", userID, err))
	}
	return &pb.SignUpResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}, nil
}
