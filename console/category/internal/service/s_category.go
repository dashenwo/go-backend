package service

import (
	"github.com/dashenwo/go-backend/v2/console/category/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// 生成验证码并发送
func (s CategoryService) Add(recipient string, recipientType int32) error {

	return nil
}
