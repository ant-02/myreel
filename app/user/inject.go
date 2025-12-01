package user

import (
	"myreel/app/user/controller/rpc"
	"myreel/app/user/domain/service"
	"myreel/app/user/infrastructure/cache"
	"myreel/app/user/infrastructure/mysql"
	"myreel/app/user/usecase"
	"myreel/config"
	"myreel/kitex_gen/user"
	"myreel/pkg/base/client"
	"myreel/pkg/constants"
	"myreel/pkg/util"
)

func InjectUserHandler() user.UserService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	sf, err := util.NewSnowflake(constants.WorkerOfUserService, config.GetDataCenterID())
	if err != nil {
		panic(err)
	}

	re, err := client.NewRedisClient(constants.RedisDBGateWay) // 使用和网关同一个数据库，目前仅用作登录登出
	if err != nil {
		panic(err)
	}

	db := mysql.NewUserDB(gormDB)
	if err := db.Magrate(); err != nil {
		panic(err)
	}
	redisCache := cache.NewUserCache(re)
	svc := service.NewUserService(db, sf, redisCache)
	uc := usecase.NewUserUseCase(db, svc, redisCache)

	return rpc.NewUserServiceImpl(uc)
}
