// Code generated by goctl. DO NOT EDIT!
// Source: history.proto

package server

import (
	"context"

	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/internal/logic"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/internal/svc"
	history "github.com/just-coding-0/learn_example/micro_service/echo/rpc/history/pb"
)

type HistoryServer struct {
	svcCtx *svc.ServiceContext
}

func NewHistoryServer(svcCtx *svc.ServiceContext) *HistoryServer {
	return &HistoryServer{
		svcCtx: svcCtx,
	}
}

func (s *HistoryServer) Add(ctx context.Context, in *history.AddRequest) (*history.AddResponse, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}

func (s *HistoryServer) Last(ctx context.Context, in *history.LastRequest) (*history.LastRequestResponse, error) {
	l := logic.NewLastLogic(ctx, s.svcCtx)
	return l.Last(in)
}
