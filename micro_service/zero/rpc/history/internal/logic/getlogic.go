package logic

import (
	"context"

	"github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/internal/svc"
	history "github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *history.GetRequest) (*history.GetResponse, error) {
	resp, err := l.svcCtx.Model.FindOne(in.Msg)
	if err != nil {
		return nil, err
	}


	return &history.GetResponse{
		Msg: resp.Msg,
		Times: resp.Times,
		LastEcho: resp.LastEcho.Format(`2006-01-02 15:04:05`),
	}, nil
}
