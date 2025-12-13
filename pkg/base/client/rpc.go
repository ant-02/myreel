package client

import (
	"errors"
	"fmt"
	"myreel/config"
	"myreel/kitex_gen/chat/chatservice"
	"myreel/kitex_gen/comment/commentservice"
	"myreel/kitex_gen/follow/followservice"
	"myreel/kitex_gen/like/likeservice"
	"myreel/kitex_gen/user/userservice"
	"myreel/kitex_gen/video/videoservice"
	"myreel/pkg/constants"

	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func initRPCClient[T any](serviceName string, newClientFunc func(string, ...client.Option) (T, error)) (*T, error) {
	if config.Etcd == nil || config.Etcd.Addr == "" {
		return nil, errors.New("config.Etcd.Addr is nil")
	}

	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		return nil, fmt.Errorf("initRPCClient etcd.NewEtcdResolver failed: %w", err)
	}

	client, err := newClientFunc(serviceName,
		client.WithResolver(r),
		client.WithMuxConnection(constants.MuxConnection),
		// client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: fmt.Sprintf(constants.KitexClientEndpointInfoFormat, serviceName)}),
	)
	if err != nil {
		return nil, fmt.Errorf("initRPCClient NewClient failed: %w", err)
	}
	return &client, nil
}

func InitUserRPC() (*userservice.Client, error) {
	return initRPCClient(constants.UserServiceName, userservice.NewClient)
}

func InitVideoRPC() (*videoservice.Client, error) {
	return initRPCClient(constants.VideoServiceName, videoservice.NewClient)
}

func InitLikeRPC() (*likeservice.Client, error) {
	return initRPCClient(string(constants.LikeServiceName), likeservice.NewClient)
}

func InitCommentRPC() (*commentservice.Client, error) {
	return initRPCClient(string(constants.CommentServiceName), commentservice.NewClient)
}

func InitFollowRPC() (*followservice.Client, error) {
	return initRPCClient(string(constants.FollowServiceName), followservice.NewClient)
}

func InitChatRPC() (*chatservice.Client, error) {
	return initRPCClient(string(constants.ChatServiceName), chatservice.NewClient)
}
