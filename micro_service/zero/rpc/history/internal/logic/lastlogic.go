package logic

import (
	"context"

	"github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/internal/svc"
	history "github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type LastLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LastLogic {
	return &LastLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LastLogic) Last(in *history.LastRequest) (*history.LastResponse, error) {
	resp, err := l.svcCtx.Model.Last()
	if err != nil {
		return nil, err
	}

	return &history.LastResponse{
		Msg: resp.Msg,
		Times: resp.Times,
		LastEcho: resp.LastEcho.Format(`2006-01-02 15:04:05`),
	}, nil
}
