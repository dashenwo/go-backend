package handler

import (
	"context"
	"encoding/json"
	conf "github.com/dashenwo/go-backend/v2/console/account/config"
	"github.com/dashenwo/go-backend/v2/console/account/global"
	"github.com/dashenwo/go-backend/v2/console/account/internal/service"
	"github.com/dashenwo/go-backend/v2/console/account/proto"
	"github.com/dashenwo/go-backend/v2/console/account/schema"
	"github.com/dashenwo/go-backend/v2/pkg/crypto"
	"github.com/dashenwo/go-backend/v2/pkg/utils/header"
	"github.com/dashenwo/go-backend/v2/pkg/utils/jwt"
	"github.com/dashenwo/go-backend/v2/pkg/utils/response/code"
	"github.com/dashenwo/go-backend/v2/pkg/utils/response/message"
	"github.com/dashenwo/go-backend/v2/pkg/utils/validate"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

type Account struct {
	accountService *service.AccountService
}

// 实例化方法
func NewAccountHandler(accountService *service.AccountService) *Account {
	return &Account{
		accountService: accountService,
	}
}

// 登录handler
func (a *Account) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}
	user, err := a.accountService.Login(req.Username, req.Password)
	if err != nil {
		return err
	}
	// 保存session
	if err = a.saveSession(ctx, user, req.Source); err != nil {
		return errors.New(conf.AppId, err.Error(), 506)
	}
	return nil
}

// 注册handler
func (a *Account) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {
	log.Info("进入了注册方法", req)
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}
	// 2.进入注册
	account, err := a.accountService.Register(req.Nickname, req.Password, req.Phone, req.Code)
	if err != nil {
		return err
	}
	res.Id = account.ID
	return nil
}

// 查询信息
func (a *Account) Info(ctx context.Context, req *proto.InfoRequest, res *proto.InfoResponse) error {
	headers := header.GetHeader(ctx)
	if req.Id == "" && headers.Get("Token") == "" {
		return errors.New(conf.AppId, message.ParameterErrorCode, code.ParameterErrorCode)
	}
	if req.Id == "" {
		data, err := jwt.Decode(headers.Get("Token"), global.Config.Session.Secret)
		if err != nil {
			return errors.New(conf.AppId, message.ParameterJwtErrorCode, code.ParameterJwtErrorCode)
		}
		req.Id = data.Id
	}

	return nil
}
func (a *Account) Update(ctx context.Context, req *proto.UpdateRequest, res *proto.UpdateResponse) error {
	return nil
}

// 保存session
func (a *Account) saveSession(ctx context.Context, data *schema.Account, source string) error {
	if data == nil {
		data = &schema.Account{}
	}
	var (
		maxAge    = global.Config.Session.Expire
		jwtData   string
		insetByte []byte
		err       error
	)
	switch source {
	case "web":
		maxAge = 3600 * 2
	case "ios":
		maxAge = 3600 * 24 * 30
	}
	claims := jwt.LoginClaims{
		Id:        data.ID,
		NickName:  data.Nickname,
		LiftTime:  maxAge,
		LoginTime: time.Now().Unix(),
	}
	jwtData, err = jwt.Encode(claims, global.Config.Session.Secret)
	if err != nil {
		return err
	}
	token := crypto.Md5(jwtData)
	insertData := make(map[string]interface{})
	insertData["token"] = token
	insertData["jwt"] = jwtData
	insertData["max_age"] = maxAge
	key := global.Config.Session.Prefix + ":" + data.ID + ":" + token
	// 序列化数据
	if insetByte, err = json.Marshal(insertData); err != nil {
		return err
	}
	// base64加密
	insetStr := crypto.Base64Encode(insetByte)
	// 插入redis
	if err = global.Redis.Set(global.Redis.Context(), key, insetStr, time.Duration(maxAge)*time.Second).Err(); err != nil {
		return err
	}
	// 发送header
	sessionId := crypto.Base64Encode([]byte(token + "|" + data.ID))
	sessionCookie := &http.Cookie{
		Name:    global.Config.Session.Name,
		Path:    "/",
		Value:   sessionId,
		Expires: time.Now().Add(time.Duration(maxAge) * time.Second),
	}
	headers := metadata.Pairs(
		"Set-Cookie", sessionCookie.String(),
	)
	_ = grpc.SendHeader(ctx, headers)
	return nil
}
