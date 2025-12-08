package like

import (
	"myreel/app/like/controller/rpc"
	"myreel/app/like/domain/service"
	"myreel/app/like/infrastructure/cache"
	"myreel/app/like/infrastructure/mysql"
	likeRpcPkg "myreel/app/like/infrastructure/rpc"
	"myreel/app/like/usecase"
	"myreel/config"
	"myreel/kitex_gen/like"
	"myreel/pkg/base/client"
	"myreel/pkg/constants"
	"myreel/pkg/util"
)

func InjectLikeHandler() like.LikeService {
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

	db := mysql.NewLikeDB(gormDB)
	if err := db.Magrate(); err != nil {
		panic(err)
	}

	videoClient, err := client.InitVideoRPC()
	if err != nil {
		panic(err)
	}
	commentClient, err := client.InitCommentRPC()
	if err != nil {
		panic(err)
	}

	lRPC := likeRpcPkg.NewLikeRpcImpl(*videoClient, *commentClient)

	redisCache := cache.NewVideoCache(re)

	svc := service.NewLikeService(db, sf, redisCache, lRPC)
	uc := usecase.NewLikeUseCase(db, svc, redisCache, lRPC)

	return rpc.NewLikeServiceImpl(uc)
}
