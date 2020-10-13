package logic

import (
	"context"
	"github.com/just-coding-0/learn_example/micro_service/zero/rpc/echo/internal/svc"
	echo "github.com/just-coding-0/learn_example/micro_service/zero/rpc/echo/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type EchoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEchoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EchoLogic {
	return &EchoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EchoLogic) Echo(in *echo.Request) (*echo.Response, error) {
	return &echo.Response{Msg: in.Msg}, nil
}
