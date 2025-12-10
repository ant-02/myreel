package pack

import (
	domainModel "myreel/app/follow/domain/model"
	"myreel/kitex_gen/follow"
)

func BuildUserProfile(f *domainModel.UserProfile) *follow.User {
	if f == nil {
		return nil
	}
	return &follow.User{
		Id:        f.Id,
		Username:  f.Username,
		AvatarUrl: f.AvatarUrl,
	}
}

func BuildUserProfiles(vs []*domainModel.UserProfile) []*follow.User {
	l := len(vs)
	if l == 0 {
		return nil
	}

	users := make([]*follow.User, l)
	for i, v := range vs {
		users[i] = BuildUserProfile(v)
	}
	return users
}

func BuildPagination(p *domainModel.Pagination) *follow.Pagination {
	return &follow.Pagination{
		NextCursor: p.NextCursor,
		PrevCursor: p.PrevCursor,
		Total:      p.Total,
	}
}

func BuildUserList(vs []*follow.User, pagination *follow.Pagination) *follow.UserList {
	return &follow.UserList{
		Items:      vs,
		Pagination: pagination,
	}
}
