package video

import (
	"myreel/app/video/controller/rpc"
	"myreel/app/video/domain/service"
	"myreel/app/video/infrastructure/cache"
	"myreel/app/video/infrastructure/mysql"
	videoRpcPkg "myreel/app/video/infrastructure/rpc"
	"myreel/app/video/usecase"
	"myreel/config"
	"myreel/kitex_gen/video"
	"myreel/pkg/base/client"
	"myreel/pkg/constants"
	"myreel/pkg/util"
)

func InjectVideoHandler() video.VideoService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	sf, err := util.NewSnowflake(constants.WorkerOfUserService, config.GetDataCenterID())
	if err != nil {
		panic(err)
	}

	re, err := client.InitRedis(constants.RedisDBGateWay) // 使用和网关同一个数据库，目前仅用作登录登出
	if err != nil {
		panic(err)
	}

	db := mysql.NewVideoDB(gormDB)
	if err := db.Magrate(); err != nil {
		panic(err)
	}

	userClient, err := client.InitUserRPC()
	if err != nil {
		panic(err)
	}

	vRPC := videoRpcPkg.NewVideoRpcImpl(*userClient)

	redisCache := cache.NewVideoCache(re)

	svc := service.NewVideoService(db, sf, redisCache, vRPC)
	uc := usecase.NewVideoUseCase(db, svc, redisCache, vRPC)

	return rpc.NewVideoServiceImpl(uc)
}
