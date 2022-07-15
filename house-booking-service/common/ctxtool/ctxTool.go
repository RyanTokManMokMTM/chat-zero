package ctxtool

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	JWTTokenUSerID = "user_id"
)

func GetUserIDFromCTX(ctx context.Context) int64 {
	var userID int64
	if jsonUerID, ok := ctx.Value(JWTTokenUSerID).(json.Number); ok {
		if uidInt64, err := jsonUerID.Int64(); err != nil {
			userID = uidInt64
		} else {
			//using for tracking
			//ctx recording
			logx.WithContext(ctx).Error(err)
		}
	}
	return userID
}
