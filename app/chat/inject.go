package chat

import (
	"myreel/app/chat/controller/rpc"
	"myreel/app/chat/domain/service"
	"myreel/app/chat/infrastructure/cache"
	"myreel/app/chat/infrastructure/mysql"
	"myreel/app/chat/usecase"
	"myreel/config"
	"myreel/kitex_gen/chat"
	"myreel/pkg/base/client"
	"myreel/pkg/constants"
	"myreel/pkg/util"
)

func InjectChatHandler() chat.ChatService {
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

	db := mysql.NewChatDB(gormDB)
	if err := db.Magrate(); err != nil {
		panic(err)
	}

	// userClient, err := client.InitUserRPC()
	// if err != nil {
	// 	panic(err)
	// }

	// vRPC := videoRpcPkg.NewVideoRpcImpl(*userClient)

	redisCache := cache.NewChatCache(re)

	svc := service.NewChatService(db, sf, redisCache, nil)
	uc := usecase.NewChatUseCase(db, svc, redisCache, nil)

	return rpc.NewChatServiceImpl(uc)
}
