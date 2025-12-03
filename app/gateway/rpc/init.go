package rpc

import (
	"myreel/kitex_gen/user/userservice"
	"myreel/kitex_gen/video/videoservice"
)

var (
	userClient  userservice.Client
	videoClient videoservice.Client
)

func Init() {
	InitUserClient()
}
