// Code generated by goctl. DO NOT EDIT!
// Source: echo.proto

//go:generate mockgen -destination ./echo_mock.go -package echo -source $GOFILE

package echo

import (
	"context"

	echo "github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/pb"

	"github.com/tal-tech/go-zero/core/jsonx"
	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Echo interface {
		Echo(ctx context.Context, in *Request) (*Response, error)
	}

	defaultEcho struct {
		cli zrpc.Client
	}
)

func NewEcho(cli zrpc.Client) Echo {
	return &defaultEcho{
		cli: cli,
	}
}

func (m *defaultEcho) Echo(ctx context.Context, in *Request) (*Response, error) {
	var request echo.Request
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := echo.NewEchoClient(m.cli.Conn())
	resp, err := client.Echo(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret Response
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}
