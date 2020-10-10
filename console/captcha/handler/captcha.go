package handler

import (
	"context"
	conf "github.com/dashenwo/go-backend/v2/console/captcha/config"
	"github.com/dashenwo/go-backend/v2/console/captcha/internal/service"
	"github.com/dashenwo/go-backend/v2/console/captcha/proto"
	"github.com/dashenwo/go-backend/v2/pkg/utils/validate"
	"time"
)

type Captcha struct {
	captchaService *service.CaptchaService
}

// 实例化方法
func NewAccountHandler(captcha *service.CaptchaService) *Captcha {
	return &Captcha{
		captchaService: captcha,
	}
}

// 登录handler
func (a *Captcha) Generate(ctx context.Context, req *proto.GenerateRequest, res *proto.GenerateResponse) error {
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}
	captcha, err := a.captchaService.Generate(req.Recipient, req.Type)
	if err != nil {
		return err
	}
	now := time.Now()
	hh, _ := time.ParseDuration("10m")
	res.Id = captcha.ID
	res.Expires = now.Add(hh).Format("2006-01-02 15:04:05")

	//userCookie := &http.Cookie{
	//	Name: "id",
	//	Value: res.Id,
	//}
	//expiresCookie := &http.Cookie{
	//	Name: "expires",
	//	Value: res.Expires,
	//}
	//header := metadata.Pairs(
	//	"Set-Cookie", userCookie.String(),
	//	"Set-Cookie", expiresCookie.String(),
	//)
	//_ = grpc.SendHeader(ctx, header)
	return nil
}

// 注册handler
func (a *Captcha) Verify(ctx context.Context, req *proto.VerifyRequest, res *proto.VerifyResponse) error {
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}
	//2.验证验证码，传入手机号或者邮箱和验证码
	_, err := a.captchaService.Verify(req.Recipient, req.Code, req.Type)
	if err != nil {
		return err
	}
	return nil
}
