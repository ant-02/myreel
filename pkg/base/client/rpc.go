package client

import (
	"fmt"
	"myreel/kitex_gen/user/userservice"
	"myreel/kitex_gen/video/videoservice"
	"myreel/pkg/constants"

	"github.com/cloudwego/kitex/client"
)

func initRPCClient[T any](serviceName string, hostPort string, newClientFunc func(string, ...client.Option) (T, error)) (*T, error) {
	client, err := newClientFunc(serviceName, client.WithHostPorts(hostPort))
	if err != nil {
		return nil, fmt.Errorf("initRPCClient NewClient failed: %w", err)
	}
	return &client, nil
}

func InitUserRPC() (*userservice.Client, error) {
	return initRPCClient(constants.UserServiceName, "127.0.0.1:20002", userservice.NewClient)
}

func InitVideoRPC() (*videoservice.Client, error) {
	return initRPCClient(constants.VideoServiceName, "127.0.0.1:20003", videoservice.NewClient)
}
