package mysql

import (
	"myreel/pkg/constants"
	"time"
)

type Like struct {
	Id        int64     `gorm:"type:bigint;primaryKey"`
	Uid       int64     `gorm:"type:bigint"`
	VideoId   int64     `gorm:"type:bigint;default:null"`
	CommentId int64     `gorm:"type:bigint;default:null"`
	Status    int64     `gorm:"type:int;default:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}

func (Like) TableName() string {
	return constants.LikeTableName
}
