package handler

import (
	"context"
	"github.com/dashenwo/go-backend/v2/console/category/internal/service"
	"github.com/dashenwo/go-backend/v2/console/category/proto"
)

type Category struct {
	categoryService *service.CategoryService
}

// 实例化方法
func NewCategoryHandler(captcha *service.CategoryService) *Category {
	return &Category{
		categoryService: captcha,
	}
}

func (c *Category) Query(ctx context.Context, req *proto.QueryRequest, res *proto.QueryResponse) error {
	panic("implement me")
}

func (c *Category) Add(ctx context.Context, req *proto.AddRequest, res *proto.AddResponse) error {
	panic("implement me")
}

func (c *Category) Edit(ctx context.Context, req *proto.EditRequest, res *proto.EditResponse) error {
	panic("implement me")
}

func (c *Category) Delete(ctx context.Context, req *proto.DeleteRequest, res *proto.DeleteResponse) error {
	panic("implement me")
}
