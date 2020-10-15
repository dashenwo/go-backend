package repository

import (
	"github.com/dashenwo/go-backend/v2/console/category/internal/model"
)

// 用户接口
type CategoryRepository interface {
	QueryOne(captcha *model.Category) (*model.Category, error)
	Insert(captcha *model.Category) error
}
