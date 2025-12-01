package client

import (
	"context"
	"fmt"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMySQL() (db *gorm.DB, err error) {
	dsn, err := util.GetMysqlDSN()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL get mysql DSN error: %v", err))
	}
	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,  // 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
			SkipDefaultTransaction: false, // 不禁用默认事务(即单个创建、更新、删除时使用事务)
			TranslateError:         true,  // 允许翻译错误
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
			Logger: glogger.New(
				logger.GetMysqlLogger(),
				glogger.Config{
					SlowThreshold:             time.Second,  // 超过一秒的查询被认为是慢查询
					LogLevel:                  glogger.Warn, // 日志等级
					IgnoreRecordNotFoundError: true,         // 当未找到(RecordNotFoundError)时候不记录
					ParameterizedQueries:      true,         // 在 SQL 中不包含参数
					Colorful:                  false,        // 禁用颜色渲染
				}),
		})
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL mysql connect error: %v", err))
	}

	sqlDB, err := db.DB() // 尝试获取 DB 实例对象
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("get generic database object error: %v", err))
	}

	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)       // 最大闲置连接数
	sqlDB.SetMaxOpenConns(constants.MaxConnections)     // 最大连接数
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime) // 最大可复用时间
	sqlDB.SetConnMaxIdleTime(constants.ConnMaxIdleTime) // 最长保持空闲状态时间
	db = db.WithContext(context.Background())

	// 进行连通性测试
	if err = sqlDB.Ping(); err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("ping database error: %v", err))
	}

	return db, nil
}
