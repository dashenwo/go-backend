package model

import (
	"github.com/dashenwo/go-backend/v2/console/category/internal/schema"
)

// 用户接口
type CategoryRepository interface {
	QueryOne(param schema.CategoryQueryOneParam) (*schema.Category, error)
	Query(param schema.CategoryQueryParam) (*schema.CategoryQueryResult, error)
	Create(param schema.Category) (*schema.CategoryCreateResult, error)
	Edit(param schema.CategoryEditParam) (*schema.CategoryEditResult, error)
	Delete(param schema.CategoryDeleteParam) (*schema.CategoryDeleteParam, error)
}
