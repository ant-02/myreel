package main

import (
	"log"
	"myreel/app/video/controller/rpc"
	video "myreel/kitex_gen/video/videoservice"
)

func main() {
	svr := video.NewServer(new(rpc.VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
