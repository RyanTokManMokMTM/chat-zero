package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	MySQL struct {
		Datasource string
	}

	JWTAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	RedisCache cache.CacheConf
	Salt       string
}
