package main

import (
	"myreel/app/like"
	"myreel/config"
	"myreel/kitex_gen/like/likeservice"
	"myreel/pkg/constants"
	"myreel/pkg/logger"
	"myreel/pkg/util"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func init() {
	config.Init(constants.LikeServiceName)
}

func main() {
	listenAddr, err := util.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Like: get available port failed, err: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("Like: resolve tcp addr failed, err: %v", err)
	}

	svr := likeservice.NewServer(
		like.InjectLikeHandler(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.Service.Name,
		}),
		server.WithServiceAddr(addr),
	)

	if err = svr.Run(); err != nil {
		logger.Fatalf("Like: run server failed, err: %v", err)
	}
}
