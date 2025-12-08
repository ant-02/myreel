package mysql

import (
	"myreel/app/follow/domain/repository"
	"myreel/pkg/errno"

	"gorm.io/gorm"
)

type followDB struct {
	client *gorm.DB
}

func NewFollowDB(client *gorm.DB) repository.FollowDB {
	return &followDB{client: client}
}

func (db *followDB) Magrate() error {
	if err := db.client.AutoMigrate(&Follow{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate follow model")
	}
	return nil
}
