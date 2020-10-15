package global

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/olivere/elastic/v7"
)

var (
	// elasticsearch客户端
	ReqClient client.Client
	EsClient  *elastic.Client
)
