package handler

import (
	"context"
	conf "github.com/dashenwo/go-backend/v2/console/category/config"
	"github.com/dashenwo/go-backend/v2/console/category/internal/schema"
	"github.com/dashenwo/go-backend/v2/console/category/internal/service"
	"github.com/dashenwo/go-backend/v2/console/category/proto"
	"github.com/dashenwo/go-backend/v2/pkg/utils/request"
	"github.com/dashenwo/go-backend/v2/pkg/utils/validate"
	"github.com/jinzhu/copier"
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
func (c *Category) Get(ctx context.Context, in *proto.QueryOneRequest, out *proto.CategorySchema) error {

	return nil
}

func (c *Category) Query(ctx context.Context, req *proto.QueryRequest, res *proto.QueryResponse) error {
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}

	res.List = nil
	return nil
}

func (c *Category) Create(ctx context.Context, req *proto.AddRequest, res *proto.AddResponse) error {
	//1.验证数据
	param := schema.Category{}
	if err := request.ShouldBind(req, &param, conf.AppId); err != nil {
		return err
	}
	data, err := c.categoryService.Create(param)
	if err != nil {
		return err
	}
	_ = copier.Copy(res, data)
	return nil
}

func (c *Category) Edit(ctx context.Context, req *proto.EditRequest, res *proto.EditResponse) error {
	panic("implement me")
}

func (c *Category) Delete(ctx context.Context, req *proto.DeleteRequest, res *proto.DeleteResponse) error {
	panic("implement me")
}
