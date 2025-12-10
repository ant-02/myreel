package rpc

import (
	"context"
	build "myreel/app/user/controller/rpc/pack"
	"myreel/app/user/usecase"
	user "myreel/kitex_gen/user"
	base "myreel/pkg/base/context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserServiceImpl(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

// Login implements the UserServiceImpl interface.
func (s *UserHandler) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	u, t, err := s.useCase.Login(ctx, req.Username, req.Password)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}
	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildUserWithToken(u, t)
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserHandler) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	if err = s.useCase.Register(ctx, req.Username, req.Password); err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserHandler) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	resp = new(user.GetUserInfoResponse)

	u, err := s.useCase.GetUserById(ctx, req.UserId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildUser(u)
	return
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserHandler) GetUploadToken(ctx context.Context, req *user.GetUploadTokenRequest) (resp *user.GetUploadTokenResponse, err error) {
	resp = new(user.GetUploadTokenResponse)

	t, err := s.useCase.GetLoadToken(ctx, req.Suffix, req.UserId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildUpyunToken(t)
	return
}

// GetMFA implements the UserServiceImpl interface.
func (s *UserHandler) GetMFA(ctx context.Context, req *user.GetMFARequest) (resp *user.GetMFAResponse, err error) {
	// TODO: Your code here...
	return
}

// BindMFA implements the UserServiceImpl interface.
func (s *UserHandler) BindMFA(ctx context.Context, req *user.BindMFARequest) (resp *user.BindMFAResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchImg implements the UserServiceImpl interface.
func (s *UserHandler) SearchImg(ctx context.Context, req *user.SearchImgRequest) (resp *user.SearchImgResponse, err error) {
	// TODO: Your code here...
	return
}

// Refresh implements the UserServiceImpl interface.
func (s *UserHandler) Refresh(ctx context.Context, req *user.RefreshRequest) (resp *user.RefreshResponse, err error) {
	resp = new(user.RefreshResponse)

	token, err := s.useCase.Refresh(ctx, req.Token, req.UserId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Token = token
	return
}

func (s *UserHandler) SetUserAvatarUrl(ctx context.Context, req *user.SetUserAvatarUrlRequest) (resp *user.SetUserAvatarUrlResponse, err error) {
	resp = new(user.SetUserAvatarUrlResponse)

	err = s.useCase.SetAvatar(ctx, req.UserId, req.Url)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	return
}

// GetUseridByUsername implements the UserServiceImpl interface.
func (s *UserHandler) GetUseridByUsername(ctx context.Context, req *user.GetUserIdByUsernameRequest) (resp *user.GetUserIdByUsernameResponse, err error) {
	resp = new(user.GetUserIdByUsernameResponse)

	uid, err := s.useCase.GetUserIdByUsername(ctx, req.Username)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.UserId = uid
	return
}

// GetUsersByIds implements the UserServiceImpl interface.
func (s *UserHandler) GetUsersByIds(ctx context.Context, req *user.GetUsersByIdsRequest) (resp *user.GetUsersByIdsResponse, err error) {
	resp = new(user.GetUsersByIdsResponse)
	users, err := s.useCase.GetUsersByIds(ctx, req.Ids)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.List = build.BuildUseProfiles(users)
	return
}
