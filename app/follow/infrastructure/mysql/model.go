package mysql

import (
	"myreel/pkg/constants"
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	Id            int64          `gorm:"type:bigint;primaryKey"`
	FolloweringId int64          `gorm:"type:bigint;not null"`
	FolloweredId  int64          `gorm:"type:bigint;not null"`
	Status        int64          `gorm:"type:int(2);noy null;default:0"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type FolloweredIdWithTime struct {
	FolloweredId int64
	CreatedAt    time.Time
}

func (Follow) TableName() string {
	return constants.FollowTableName
}
