package rpc

import "myreel/kitex_gen/user/userservice"

var userClient userservice.Client

func Init() {
	InitUserClient()
}
