// Code generated by goctl. DO NOT EDIT!
// Source: echo.proto

package main

import (
	"flag"
	"fmt"

	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/internal/config"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/internal/server"
	"github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/internal/svc"
	echo "github.com/just-coding-0/learn_example/micro_service/echo/rpc/echo/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/echo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	echoSrv := server.NewEchoServer(ctx)

	s, err := zrpc.NewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		echo.RegisterEchoServer(grpcServer, echoSrv)
	})
	logx.Must(err)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
