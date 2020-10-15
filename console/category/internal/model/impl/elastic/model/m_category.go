package model

import (
	"context"
	"github.com/dashenwo/go-backend/v2/console/category/global"
	"github.com/dashenwo/go-backend/v2/console/category/internal/model"
	"github.com/dashenwo/go-backend/v2/console/category/internal/model/impl/elastic/entity"
	"github.com/dashenwo/go-backend/v2/console/category/internal/schema"
	"github.com/dashenwo/go-backend/v2/pkg/utils/generate"
	"strconv"
)

type CaptchaModel struct {
}

func (c CaptchaModel) QueryOne(param schema.CategoryQueryOneParam) (*schema.Category, error) {
	panic("implement me")
}

func (c CaptchaModel) Query(param schema.CategoryQueryParam) (*schema.CategoryQueryResult, error) {
	panic("implement me")
}

func (c CaptchaModel) Create(param schema.Category) (*schema.CategoryCreateResult, error) {
	category := entity.SchemaCategory(param).ToCategory()
	if err := c.CheckUrl(category); err != nil {
		return nil, err
	}
	// 设置id
	callRes, callErr := generate.GetSnowflakeId(global.ReqClient)
	if callErr != nil {
		return nil, callErr
	}
	category.Id = strconv.FormatInt(callRes.Id, 10)
	res, err := global.EsClient.Index().
		Index("db_category").
		Id(category.Id).
		BodyJson(category).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &schema.CategoryCreateResult{
		Id: res.Id,
	}, nil
}

func (c CaptchaModel) Edit(param schema.CategoryEditParam) (*schema.CategoryEditResult, error) {
	panic("implement me")
}

func (c CaptchaModel) Delete(param schema.CategoryDeleteParam) (*schema.CategoryDeleteParam, error) {
	panic("implement me")
}

func (c CaptchaModel) CheckUrl(category *entity.Category) error {
	return nil
}

func NewCaptchaModel() model.CategoryRepository {
	return &CaptchaModel{}
}
