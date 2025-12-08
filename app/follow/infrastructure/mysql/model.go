package mysql

import (
	"myreel/pkg/constants"
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	Id          string         `gorm:"type:varchar(100);primaryKey"`
	FollowingId string         `gorm:"type:varchar(100);not null"`
	FollowerId  string         `gorm:"type:varchar(100);not null"`
	Status      int64          `gorm:"type:int(2);noy null;default:0"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Follow) TableName() string {
	return constants.FollowTableName
}
