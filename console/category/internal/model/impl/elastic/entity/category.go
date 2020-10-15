package entity

import (
	"github.com/dashenwo/go-backend/v2/console/category/internal/schema"
	"github.com/jinzhu/copier"
	"time"
)

type Category struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Url         string   `json:"url"`
	Tags        []string `json:"tags"`
	Sort        int      `json:"sort"`
	CreatedTime int64    `json:"created_time"`
	UpdatedTime int64    `json:"updated_time"`
}

func (c Category) ToProtoCategory() *schema.Category {
	item := new(schema.Category)
	_ = copier.Copy(item, c)
	return item
}

//---------Categorys----------
type CategoryList []*Category

func (c CategoryList) ToProtoCategory() []*schema.Category {
	list := make([]*schema.Category, len(c))
	for i, item := range c {
		list[i] = item.ToProtoCategory()
	}
	return list
}

type SchemaCategory schema.Category

func (c SchemaCategory) ToCategory() *Category {
	item := new(Category)
	_ = copier.Copy(item, c)
	// 设置默认创建时间
	now := time.Now().Unix()
	if item.CreatedTime == 0 {
		item.CreatedTime = now
	}
	// 设置默认修改时间
	item.UpdatedTime = now
	return item
}
