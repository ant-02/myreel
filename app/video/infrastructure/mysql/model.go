package mysql

import (
	"myreel/pkg/constants"
	"time"
)

type Video struct {
	Id           int64     `gorm:"type:bigint;primaryKey"`
	Uid          int64     `gorm:"type:bigint"`
	Title        string    `gorm:"type:varchar(100);not null"`
	Description  string    `gorm:"type:varchar(256);not null"`
	VideoUrl     string    `gorm:"type:varchar(256);unique;not null"`
	CoverUrl     string    `gorm:"type:varchar(256)"`
	VisitCount   int64     `gorm:"type:int;default:0"`
	LikeCount    int64     `gorm:"type:int;default:0"`
	CommentCount int64     `gorm:"type:int;default:0"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    time.Time `gorm:"type:datetime;default:null"`
}

func (Video) TableName() string {
	return constants.VideoTableName
}
