package global

import (
	"github.com/bwmarrin/snowflake"
	"github.com/micro/go-micro/v2/client"
)

var (
	SnowflakeNode *snowflake.Node
	ReqClient     client.Client
)
