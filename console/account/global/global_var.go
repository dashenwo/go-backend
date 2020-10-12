package global

import "github.com/dashenwo/go-backend/v2/console/account/config"

var (
	// 配置信息
	Config = &config.Config{
		Session: config.Session{
			Prefix: "session",
			Name:   "sessionid",
			Expire: 3600,
			Secret: "",
		},
	}
)
