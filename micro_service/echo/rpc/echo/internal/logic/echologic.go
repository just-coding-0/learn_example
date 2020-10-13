package logic

import (
	"context"
	"fmt"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/internal/svc"
	echo "github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/pb"

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

// 填充逻辑
func (l *EchoLogic) Echo(in *echo.Request) (*echo.Response, error) {
	return &echo.Response{Msg: fmt.Sprintf("%s",in.Msg)}, nil
}
