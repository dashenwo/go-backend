package registry

import (
	"github.com/dashenwo/go-backend/v2/console/account/global"
	"github.com/dashenwo/go-backend/v2/console/account/internal/repository/persistence/gorm"
	"github.com/dashenwo/go-backend/v2/console/account/internal/service"
	"github.com/dashenwo/go-backend/v2/pkg/storage/elasticsearch"
	"github.com/dashenwo/go-backend/v2/pkg/storage/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	buildProvide(c)
	return c, nil
}

func buildProvide(c *dig.Container) {
	var err error
	// 初始化配置信息
	if err := config.Scan(global.Config); err != nil {
		panic("初始化配置信息失败")
	}
	// DB初始化
	gorm.InitDb()
	// 初始化elasticsearch
	if es, err := elasticsearch.Init(global.Config.Elasticsearch); err == nil {
		global.Es = es
	} else {
		panic("初始化es失败")
	}
	// 初始化redis
	if global.Redis, err = redis.Init(&global.Config.Redis); err != nil {
		panic("redis初始化失败" + err.Error())
	}

	err2 := c.Provide(gorm.NewAccountRepository)
	log.Info(err2)
	err3 := c.Provide(service.NewAccountService)
	log.Info(err3)
}
