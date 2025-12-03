package model

type Video struct {
	Id           int64
	Uid          int64
	Title        string
	Description  string
	VideoUrl     string
	CoverUrl     string
	VisitCount   int64
	LikeCount    int64
	CommentCount int64
}
