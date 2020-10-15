package config

import (
	"github.com/dashenwo/go-backend/v2/pkg/storage/elasticsearch"
	"github.com/dashenwo/go-backend/v2/pkg/storage/redis"
)

var (
	ConfPath = "./conf/"
	AppId    = "com.dashenwo.srv.category"
)

type Config struct {
	Elasticsearch elasticsearch.Elasticsearch
	Redis         redis.Redis
}
