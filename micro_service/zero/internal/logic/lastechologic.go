package logic

import (
	"context"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/svc"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/types"
	"github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/history"
	"github.com/tal-tech/go-zero/core/logx"
)

type LastEchoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLastEchoLogic(ctx context.Context, svcCtx *svc.ServiceContext) LastEchoLogic {
	return LastEchoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LastEchoLogic) LastEcho() (*types.LastEchoResponse, error) {
	resp, err := l.svcCtx.History.Last(l.ctx, &history.LastRequest{})
	if err != nil {
		return nil, err
	}
	return &types.LastEchoResponse{
		LastEcho: resp.LastEcho,
		Times:    resp.Times,
		Msg:      resp.Msg,
	}, nil
}
