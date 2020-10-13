package logic

import (
	"context"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/internal/svc"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/model"
	history "github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/pb"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 填充逻辑
func (l *AddLogic) Add(in *history.AddRequest) (*history.AddResponse, error) {
	_, err := l.svcCtx.Model.Insert(model.History{
		Msg:      in.Msg,
		Times:    1,
		LastEcho: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &history.AddResponse{Ok: true}, nil
}
