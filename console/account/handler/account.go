package handler

import (
	"context"
	"encoding/json"
	conf "github.com/dashenwo/go-backend/v2/console/account/config"
	"github.com/dashenwo/go-backend/v2/console/account/global"
	"github.com/dashenwo/go-backend/v2/console/account/internal/service"
	"github.com/dashenwo/go-backend/v2/console/account/proto"
	"github.com/dashenwo/go-backend/v2/pkg/crypto"
	"github.com/dashenwo/go-backend/v2/pkg/utils/validate"
	"github.com/fatih/structs"
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
	if err = a.saveSession(ctx, structs.Map(user), user.ID, req.Source); err != nil {
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
	return nil
}
func (a *Account) Update(ctx context.Context, req *proto.UpdateRequest, res *proto.UpdateResponse) error {
	return nil
}

// 保存session
func (a *Account) saveSession(ctx context.Context, data map[string]interface{}, userId, source string) error {
	if data == nil {
		data = make(map[string]interface{})
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
	jwtData, err = crypto.CreateToken(userId, global.Config.Session.Secret, maxAge)
	if err != nil {
		return err
	}
	token := crypto.Md5(jwtData)
	data["token"] = token
	key := global.Config.Session.Prefix + ":" + token
	// 序列化数据
	if insetByte, err = json.Marshal(data); err != nil {
		return err
	}
	insetStr := crypto.Base64Encode(insetByte)
	// 插入redis
	if err = global.Redis.Set(global.Redis.Context(), key, insetStr, time.Duration(maxAge)*time.Second).Err(); err != nil {
		return err
	}
	// 发送header
	sessionCookie := &http.Cookie{
		Name:    global.Config.Session.Name,
		Path:    "/",
		Value:   token,
		Expires: time.Now().Add(time.Duration(maxAge) * time.Second),
	}
	header := metadata.Pairs(
		"Set-Cookie", sessionCookie.String(),
	)
	_ = grpc.SendHeader(ctx, header)
	return nil
}
