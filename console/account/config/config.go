package config

import (
	"github.com/dashenwo/go-backend/v2/pkg/storage/elasticsearch"
	"github.com/dashenwo/go-backend/v2/pkg/storage/redis"
	"time"
)

var (
	ConfPath = "./conf/"
	AppId    = "com.dashenwo.srv.account"
)

type Config struct {
	// 数据库配置信息
	Database Database
	// session配置
	Session Session
	// es配置信息
	Elasticsearch elasticsearch.Elasticsearch
	// redis配置
	Redis redis.Redis
}

// 数据库配置信息
type Database struct {
	LogMode     bool
	AutoMigrate bool
	Engine      string
	Host        string
	Port        string
	User        string
	Password    string
	Name        string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type Session struct {
	Prefix string
	// session名
	Name string
	// 默认的过期时间
	Expire int64
	// 密匙
	Secret string
}
