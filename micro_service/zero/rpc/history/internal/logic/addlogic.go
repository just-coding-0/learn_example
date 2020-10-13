package logic

import (
	"context"
	"github.com/just-coding-0/learn_example/micro_service/zero/rpc/model"
	"time"

	"github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/internal/svc"
	history "github.com/just-coding-0/learn_example/micro_service/zero/rpc/history/pb"

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

func (l *AddLogic) Add(in *history.AddRequest) (*history.AddResponse, error) {
	_, err := l.svcCtx.Model.Insert(model.History{
		Msg:      in.Msg,
		LastEcho: time.Now(),
		Times:    1,
	})
	if err != nil {
		return nil, err
	}

	return &history.AddResponse{Ok: true}, nil
}
