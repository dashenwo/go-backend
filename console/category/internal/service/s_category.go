package service

import (
	"github.com/dashenwo/go-backend/v2/console/category/internal/model"
	"github.com/dashenwo/go-backend/v2/console/category/internal/schema"
)

type CategoryService struct {
	repo model.CategoryRepository
}

func NewCategoryService(repo model.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// 新建分类
func (s CategoryService) Create(param schema.Category) (*schema.CategoryCreateResult, error) {
	return s.repo.Create(param)
}
