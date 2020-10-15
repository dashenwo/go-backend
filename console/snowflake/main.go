package main

import (
	conf "github.com/dashenwo/go-backend/v2/console/snowflake/config"
	"github.com/dashenwo/go-backend/v2/console/snowflake/global"
	"github.com/dashenwo/go-backend/v2/console/snowflake/handler"
	"github.com/dashenwo/go-backend/v2/console/snowflake/registry"
	tracer "github.com/dashenwo/go-backend/v2/pkg/opentracing"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/logger/zap/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	zap2 "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	// 初始化日志库
	encodingConfig := zap2.NewProductionEncoderConfig()
	// 时间格式化
	encodingConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	l, err := zap.NewLogger(
		zap.WithCallerSkip(4),
		zap.WithConfig(zap2.NewProductionConfig()),
		zap.WithEncoderConfig(encodingConfig),
	)
	if err != nil {
		log.Fatal(err)
	}
	logger.DefaultLogger = l

	md := make(map[string]string)
	md["chain"] = "gray"

	gsrv := grpc.NewServer()
	//创建服务
	service := micro.NewService(
		micro.Server(gsrv),
		micro.Name(conf.AppId),
		micro.Version("latest"),
		micro.Metadata(md),
		// 设置启动ip
		micro.Address(":8001"),
		micro.Flags(
			&cli.StringFlag{
				Name:  "conf_path",
				Value: "./conf/",
				Usage: "配置文件目录",
			},
		),
		micro.Action(func(ctx *cli.Context) error {
			confPath := ctx.String("conf_path")
			conf.ConfPath = confPath
			// 配置加载
			err := config.LoadFile(conf.ConfPath + "config.yaml")
			return err
		}),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer(conf.AppId, "203.195.200.40:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	service.Init(
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		micro.WrapCall(opentracing.NewCallWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)
	// 初始化服务
	service.Init()

	// 设置grpc客户端
	global.ReqClient = service.Client()

	// Register Handler
	c, err := registry.NewContainer(service.Server())
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	// Register Handler
	handler.Apply(service.Server(), c)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
