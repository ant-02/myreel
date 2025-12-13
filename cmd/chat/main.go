package main

import (
	"myreel/app/chat"
	"myreel/config"
	"myreel/kitex_gen/chat/chatservice"
	"myreel/pkg/constants"
	"myreel/pkg/logger"
	"myreel/pkg/util"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func init() {
	config.Init(constants.ChatServiceName)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		logger.Fatalf("Payment: new etcd registry failed, err: %v", err)
	}

	listenAddr, err := util.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Video: get available port failed, err: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("Video: resolve tcp addr failed, err: %v", err)
	}

	svr := chatservice.NewServer(
		chat.InjectChatHandler(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.Service.Name,
		}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	if err = svr.Run(); err != nil {
		logger.Fatalf("Chat: run server failed, err: %v", err)
	}
}
