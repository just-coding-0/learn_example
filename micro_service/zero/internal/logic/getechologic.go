package logic

import (
	"context"
	"github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/history"

	"github.com/just-coding-0/learn_example/micro_service/zero/internal/svc"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetEchoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEchoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetEchoLogic {
	return GetEchoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEchoLogic) GetEcho(req types.GetEchoStatsRequest) (*types.GetEchoStatsResponse, error) {
	resp,err:=l.svcCtx.History.Get(l.ctx,&history.GetRequest{
		Msg: req.Msg,
	})
	if err != nil {
		return nil,err
	}

	return &types.GetEchoStatsResponse{
		LastEcho: resp.LastEcho,
		Times:    resp.Times,
		Msg:      resp.Msg,
	}, nil
}
