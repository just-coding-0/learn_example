package logic

import (
	"context"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/internal/svc"
	history "github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/pb"

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

func (l *LastLogic) Last(in *history.LastRequest) (*history.LastRequestResponse, error) {
	result, err := l.svcCtx.Model.Last()

	if err != nil {
		return nil, err
	}
	return &history.LastRequestResponse{
		Msg: result.Msg,
		Times: result.Times,
		LastEcho: result.LastEcho.Format(`2006-01-02 15:04:05`),
	}, nil
}
