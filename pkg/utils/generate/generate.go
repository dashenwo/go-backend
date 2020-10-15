package generate

import (
	"context"
	"github.com/dashenwo/go-backend/v2/console/snowflake/proto"
	"github.com/micro/go-micro/v2/client"
)

func GetSnowflakeId(client client.Client) (*proto.Response, error) {
	// create the proto client for helloworld
	srv := proto.NewSnowflakeService("com.dashenwo.srv.snowflake", client)
	// call an endpoint on the service
	rsp, callErr := srv.Generate(context.Background(), &proto.Request{})
	if callErr != nil {
		return nil, callErr
	}
	return rsp, nil
}
