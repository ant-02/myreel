package rpc

import (
	"myreel/kitex_gen/comment/commentservice"
	"myreel/kitex_gen/like/likeservice"
	"myreel/kitex_gen/user/userservice"
	"myreel/kitex_gen/video/videoservice"
)

var (
	userClient    userservice.Client
	videoClient   videoservice.Client
	likeClient    likeservice.Client
	commentClient commentservice.Client
)

func Init() {
	InitUserClient()
	InitVideoClient()
	InitLikeClient()
	InitCommentClient()
}
