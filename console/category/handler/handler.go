package handler

import (
	"github.com/dashenwo/go-backend/v2/console/category/internal/service"
	"github.com/dashenwo/go-backend/v2/console/category/proto"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/dig"
)

func Apply(server server.Server, c *dig.Container) {
	_ = c.Invoke(func(category *service.CategoryService) {
		_ = proto.RegisterCategoryHandler(server, NewCategoryHandler(category))
	})
}
