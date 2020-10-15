package registry

import (
	"github.com/dashenwo/go-backend/v2/console/category/global"
	"github.com/dashenwo/go-backend/v2/console/category/internal/repository/persistence/elastic"
	"github.com/dashenwo/go-backend/v2/console/category/internal/service"
	"github.com/dashenwo/go-backend/v2/pkg/storage/elasticsearch"
	"github.com/micro/go-micro/v2/config"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	return c, buildProvide(c)
}

func buildProvide(c *dig.Container) error {
	var err error
	// 初始化配置文件
	if err = config.Scan(global.Config); err != nil {
		return err
	}
	// 初始化es客户端
	if global.EsClient, err = elasticsearch.Init(global.Config.Elasticsearch); err != nil {
		return err
	}

	_ = c.Provide(elastic.NewCaptchaRepository)
	_ = c.Provide(service.NewCategoryService)

	return nil
}
