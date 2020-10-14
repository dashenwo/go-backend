package header

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type headerKey struct{}

func NewHeaderWrapper() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			// 把metadata转换为http的header
			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				header := http.Header{}
				for key, value := range md {
					header.Add(key, value[0])
				}
				ctx = context.WithValue(ctx, headerKey{}, header)
			}
			return h(ctx, req, rsp)
		}
	}
}

func GetHeader(ctx context.Context) http.Header {
	return ctx.Value(headerKey{}).(http.Header)
}
