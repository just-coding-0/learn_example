package svc

import (
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/internal/config"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c     config.Config
	Model *model.HistoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewHistoryModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
	}
}
