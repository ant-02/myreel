package follow

import (
	"myreel/app/follow/controller/rpc"
	"myreel/app/follow/domain/service"
	"myreel/app/follow/infrastructure/mysql"
	followRpcPkg "myreel/app/follow/infrastructure/rpc"
	"myreel/app/follow/usecase"
	"myreel/config"
	"myreel/kitex_gen/follow"
	"myreel/pkg/base/client"
	"myreel/pkg/constants"
	"myreel/pkg/util"
)

func InjectFollowHandler() follow.FollowService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	sf, err := util.NewSnowflake(constants.WorkerOfUserService, config.GetDataCenterID())
	if err != nil {
		panic(err)
	}

	// re, err := client.InitRedis(constants.RedisDBGateWay) // 使用和网关同一个数据库，目前仅用作登录登出
	// if err != nil {
	// 	panic(err)
	// }

	userClient, err := client.InitUserRPC()
	if err != nil {
		panic(err)
	}

	fRPC := followRpcPkg.NewFollowRpcImpl(*userClient)

	db := mysql.NewFollowDB(gormDB)
	if err := db.Magrate(); err != nil {
		panic(err)
	}

	// redisCache := cache.NewVideoCache(re)

	svc := service.NewFollowService(db, sf, fRPC)
	uc := usecase.NewFollowUseCase(db, svc, fRPC)

	return rpc.NewFollowServiceImpl(uc)
}
