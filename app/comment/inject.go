package comment

import (
	"myreel/app/comment/controller/rpc"
	"myreel/app/comment/domain/service"
	"myreel/app/comment/infrastructure/mysql"
	commentRpcPkg "myreel/app/comment/infrastructure/rpc"
	"myreel/app/comment/usecase"
	"myreel/config"
	"myreel/kitex_gen/comment"
	"myreel/pkg/base/client"
	"myreel/pkg/constants"
	"myreel/pkg/util"
)

func InjectCommentHandler() comment.CommentService {
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

	db := mysql.NewCommentDB(gormDB)
	if err := db.Magrate(); err != nil {
		panic(err)
	}

	videoClient, err := client.InitVideoRPC()
	if err != nil {
		panic(err)
	}

	vRpc := commentRpcPkg.NewCommentRpcImpl(*videoClient)

	svc := service.NewCommentService(db, sf, vRpc)
	uc := usecase.NewCommentUseCase(db, svc, vRpc)

	return rpc.NewCommentServiceImpl(uc)
}
