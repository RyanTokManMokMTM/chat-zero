package ctxtool

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var CTXJWTUserID = "user_id"

func GetUserIDFromCtx(ctx context.Context) uint {
	var userID uint
	if jwtID, ok := ctx.Value(CTXJWTUserID).(json.Number); ok {
		if id, err := jwtID.Int64(); err == nil {
			userID = uint(id)
		} else {
			logx.WithContext(ctx).Error(err)
		}
	}

	return userID
}
