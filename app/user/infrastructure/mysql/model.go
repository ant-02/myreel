package mysql

import (
	"myreel/pkg/constants"
	"time"
)

type User struct {
	Id        int64     `gorm:"type:bigint;primaryKey"`
	Username  string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	AvatarUrl string    `gorm:"type:varchar(256)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}

func (User) TableName() string {
	return constants.UserTableName
}
