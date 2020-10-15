package registry

import (
	"context"
	conf "github.com/dashenwo/go-backend/v2/console/category/config"
	"github.com/dashenwo/go-backend/v2/console/category/global"
	"github.com/dashenwo/go-backend/v2/console/category/internal/model/impl/elastic/model"
	"github.com/dashenwo/go-backend/v2/console/category/internal/service"
	"github.com/dashenwo/go-backend/v2/pkg/storage/elasticsearch"
	"github.com/micro/go-micro/v2/config"
	"go.uber.org/dig"
	"io/ioutil"
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
	// 创建索引
	exists, err := global.EsClient.IndexExists("db_category").Do(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		b, err := ioutil.ReadFile(conf.ConfPath + "mapping.json")
		if err != nil {
			return err
		}
		createIndex, err := global.EsClient.CreateIndex("db_category").BodyString(string(b)).Do(context.Background())
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {

		}
	}
	_ = c.Provide(model.NewCaptchaModel)
	_ = c.Provide(service.NewCategoryService)

	return nil
}
