package main

import (
	"myreel/app/user"
	"myreel/config"
	"myreel/kitex_gen/user/userservice"
	"myreel/pkg/constants"
	"myreel/pkg/logger"
	"myreel/pkg/util"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func init() {
	config.Init(constants.UserServiceName)
}

func main() {
	listenAddr, err := util.GetAvailablePort()
	if err != nil {
		logger.Fatalf("User: get available port failed, err: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("User: resolve tcp addr failed, err: %v", err)
	}

	svr := userservice.NewServer(
		user.InjectUserHandler(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.Service.Name,
		}),
		server.WithServiceAddr(addr),
	)

	if err = svr.Run(); err != nil {
		logger.Fatalf("User: run server failed, err: %v", err)
	}
}
