package main

import (
	"context"
	user "myreel/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (resp *user.UploadAvatarResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMFA implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetMFA(ctx context.Context, req *user.GetMFARequest) (resp *user.GetMFAResponse, err error) {
	// TODO: Your code here...
	return
}

// BindMFA implements the UserServiceImpl interface.
func (s *UserServiceImpl) BindMFA(ctx context.Context, req *user.BindMFARequest) (resp *user.BindMFAResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchImg implements the UserServiceImpl interface.
func (s *UserServiceImpl) SearchImg(ctx context.Context, req *user.SearchImgRequest) (resp *user.SearchImgResponse, err error) {
	// TODO: Your code here...
	return
}

// Refresh implements the UserServiceImpl interface.
func (s *UserServiceImpl) Refresh(ctx context.Context, req *user.RefreshRequest) (resp *user.RefreshResponse, err error) {
	// TODO: Your code here...
	return
}
