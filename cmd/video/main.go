package main

import (
	"myreel/app/video"
	"myreel/config"
	"myreel/kitex_gen/video/videoservice"
	"myreel/pkg/constants"
	"myreel/pkg/logger"
	"myreel/pkg/util"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func init() {
	config.Init(constants.VideoServiceName)
}

func main() {
	listenAddr, err := util.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Video: get available port failed, err: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("Video: resolve tcp addr failed, err: %v", err)
	}

	svr := videoservice.NewServer(
		video.InjectVideoHandler(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.Service.Name,
		}),
		server.WithServiceAddr(addr),
	)

	if err = svr.Run(); err != nil {
		logger.Fatalf("Video: run server failed, err: %v", err)
	}
}
