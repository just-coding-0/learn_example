package svc

import (
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/internal/config"
)

type ServiceContext struct {
	c config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
	}
}
