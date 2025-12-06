package main

import (
	"myreel/app/comment"
	"myreel/config"
	"myreel/kitex_gen/comment/commentservice"
	"myreel/pkg/constants"
	"myreel/pkg/logger"
	"myreel/pkg/util"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func init() {
	config.Init(constants.CommentServiceName)
}

func main() {
	listenAddr, err := util.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Comment: get available port failed, err: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("Comment: resolve tcp addr failed, err: %v", err)
	}

	svr := commentservice.NewServer(
		comment.InjectCommentHandler(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.Service.Name,
		}),
		server.WithServiceAddr(addr),
	)

	if err = svr.Run(); err != nil {
		logger.Fatalf("Comment: run server failed, err: %v", err)
	}
}
