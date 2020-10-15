package elastic

import (
	"github.com/dashenwo/go-backend/v2/console/category/internal/model"
	"github.com/dashenwo/go-backend/v2/console/category/internal/repository"
)

type CategoryRepository struct {
}

func (c *CategoryRepository) QueryOne(captcha *model.Category) (*model.Category, error) {
	panic("implement me")
}

func (c *CategoryRepository) Insert(captcha *model.Category) error {
	panic("implement me")
}

func NewCaptchaRepository() repository.CategoryRepository {
	return &CategoryRepository{}
}
