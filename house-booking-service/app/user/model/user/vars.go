package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var (
	AuthTypeSystem  = "system" //login as system service
	AuthTypeSmallWX = "wxMini" //login as WeChat mini
)
