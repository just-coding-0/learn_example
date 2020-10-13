package svc

import (
	"github.com/just-coding-0/learn_example/micro_service/echo/internal/config"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/echo"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/history"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Echo    echo.Echo
	History history.History
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Echo:    echo.NewEcho(zrpc.MustNewClient(c.Echo)),
		History: history.NewHistory(zrpc.MustNewClient(c.History)),
	}
}
