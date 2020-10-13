package logic

import (
	"context"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/echo"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/history"

	"github.com/just-coding-0/learn_example/micro_service/echo/internal/svc"
	"github.com/just-coding-0/learn_example/micro_service/echo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EchoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEchoLogic(ctx context.Context, svcCtx *svc.ServiceContext) EchoLogic {
	return EchoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 填充逻辑
func (l *EchoLogic) Echo(req types.EchoRequest) (*types.EchoResponse, error) {
	resp, err := l.svcCtx.Echo.Echo(l.ctx, &echo.Request{Msg: req.Msg})
	if err != nil {
		return nil, err
	}

	_resp, err := l.svcCtx.History.Add(l.ctx, &history.AddRequest{Msg: resp.Msg})
	if err != nil {
		l.Logger.Errorf("History.Add: error %v", err)
	}
	if !_resp.Ok {
		l.Logger.Error("History.Add: failure ")
	}

	return &types.EchoResponse{Msg: resp.Msg}, nil
}
