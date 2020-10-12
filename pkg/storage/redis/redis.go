package redis

import (
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type Redis struct {
	// redis地址
	Hosts []string
	// 密码
	Password string
	// 超时时间
	TimeOut time.Duration
}

var (
	client redis.UniversalClient
	once   sync.Once
	err    error
)

func Init(conf *Redis) (redis.UniversalClient, error) {
	once.Do(func() {
		// 单节点
		if len(conf.Hosts) > 1 {
			options := redis.ClusterOptions{
				Addrs: conf.Hosts,
			}
			if conf.Password != "" {
				options.Password = conf.Password
			}
			client = redis.NewClusterClient(&options)
		} else {
			options := redis.Options{
				Addr: conf.Hosts[0],
			}
			if conf.Password != "" {
				options.Password = conf.Password
			}
			client = redis.NewClient(&options)
		}
		err = client.Ping(client.Context()).Err()
	})
	return client, err
}
