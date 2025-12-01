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
	resp.Data.User = build.BuildUser(u)
	resp.Data.Token = build.BuildToken(t)
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
	// TODO: Your code here...
	return
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserHandler) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (resp *user.UploadAvatarResponse, err error) {
	// TODO: Your code here...
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
	// TODO: Your code here...
	return
}
