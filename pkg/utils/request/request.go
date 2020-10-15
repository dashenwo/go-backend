package request

import (
	"github.com/dashenwo/go-backend/v2/pkg/utils/validate"
	"github.com/jinzhu/copier"
)

func ShouldBind(req interface{}, param interface{}, appId ...string) error {
	id := "dashenwo.comm.go-library.micro.ShouldBind"
	if len(appId) > 1 {
		id = appId[0]
	}
	//1.验证数据
	if err := validate.Validate(req, id); err != nil {
		return err
	}
	_ = copier.Copy(param, req)
	return nil
}
