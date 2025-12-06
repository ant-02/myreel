package model

type Comment struct {
	Id         int64
	VideoId    int64
	Uid        int64
	ParentId   int64
	LikeCount  int64
	ChildCount int64
	Content    string
	CreatedAt  int64
	UpdatedAt  int64
}

type Pagination struct {
	NextCursor int64
	PrevCursor int64
	Total      int64
}
