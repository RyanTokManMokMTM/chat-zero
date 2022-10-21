package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	MySQL struct {
		Datasource string
		//SetConnMaxIdleTime
		//SetMaxIdleConns
		//SetMaxOpenConns
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Salt string
}
