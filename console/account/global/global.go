package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2/client"
	"github.com/olivere/elastic/v7"
)

var (
	// elasticsearch客户端
	Es        *elastic.Client
	Redis     redis.UniversalClient
	ReqClient client.Client
)
