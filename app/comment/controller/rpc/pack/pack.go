package pack

import (
	domainModel "myreel/app/comment/domain/model"
	"myreel/kitex_gen/comment"
)

func BuildComment(c *domainModel.Comment) *comment.Comment {
	if c == nil {
		return nil
	}
	var lc, cc int64
	lc, cc = c.LikeCount, c.ChildCount
	return &comment.Comment{
		Id:         c.Id,
		Uid:        c.Uid,
		VideoId:    c.VideoId,
		ParentId:   c.ParentId,
		LikeCount:  &lc,
		ChildCount: &cc,
		Content:    c.Content,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func BuildComments(cs []*domainModel.Comment) []*comment.Comment {
	if cs == nil {
		return nil
	}

	l := len(cs)
	comments := make([]*comment.Comment, l)
	for i, c := range cs {
		comments[i] = BuildComment(c)
	}
	return comments
}

func BuildPagination(p *domainModel.Pagination) *comment.Pagination {
	return &comment.Pagination{
		NextCursor: p.NextCursor,
		PrevCursor: p.PrevCursor,
		Total:      p.Total,
	}
}

func BuildCommentList(cs []*comment.Comment, p *comment.Pagination) *comment.CommentList {
	return &comment.CommentList{
		Items:      cs,
		Pagination: p,
	}
}
