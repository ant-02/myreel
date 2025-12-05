package main

import (
	"log"
	"myreel/app/like/controller/rpc"
	like "myreel/kitex_gen/like/likeservice"
)

func main() {
	svr := like.NewServer(new(rpc.LikeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
