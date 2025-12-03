package rpc

import (
	"context"
	api "myreel/app/gateway/model/api/user"
	"myreel/kitex_gen/user"
	"myreel/kitex_gen/user/userservice"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"

	"github.com/cloudwego/kitex/client"
)

func InitUserClient() {
	c, err := userservice.NewClient(constants.UserServiceName, client.WithHostPorts("0.0.0.0:20002"))
	if err != nil {
		logger.Fatalf("api.rpc.user InitUserRPC failed, err is %v", err)
	}
	userClient = c
}

func RegisterRPC(ctx context.Context, req *user.RegisterRequest) error {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		logger.Errorf("RegisterRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return nil
}

func LoginRPC(ctx context.Context, req *user.LoginRequest) (*api.LoginResponse, error) {
	resp, err := userClient.Login(ctx, req)
	if err != nil {
		logger.Errorf("LoginRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}

	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return &api.LoginResponse{
		User: &api.User{
			Id:        resp.Data.User.Id,
			Username:  resp.Data.User.Username,
			Password:  resp.Data.User.Password,
			AvatarUrl: resp.Data.User.AvatarUrl,
		},
		Token: &api.Token{
			AccessToken:       resp.Data.Token.AccessToken,
			AccessExpireTime:  resp.Data.Token.AccessExpireTime,
			RefreshToken:      resp.Data.Token.RefreshToken,
			RefreshExpireTime: resp.Data.Token.RefreshExpireTime,
		},
	}, nil
}

func GetUserByIdRPC(ctx context.Context, req *user.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	resp, err := userClient.GetUserInfo(ctx, req)
	if err != nil {
		logger.Errorf("GetUserByIdRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}

	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return &api.GetUserInfoResponse{
		User: &api.User{
			Id:        resp.Data.Id,
			Username:  resp.Data.Username,
			Password:  resp.Data.Password,
			AvatarUrl: resp.Data.AvatarUrl,
		},
	}, nil
}

func RefreshRPC(ctx context.Context, req *user.RefreshRequest) (*api.RefreshResponse, error) {
	resp, err := userClient.Refresh(ctx, req)
	if err != nil {
		logger.Errorf("RefreshRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}

	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return &api.RefreshResponse{
		Token: resp.Token,
	}, nil
}

func GetUploadTokenRPC(ctx context.Context, req *user.GetUploadTokenRequest) (*api.GetUploadTokenResponse, error) {
	resp, err := userClient.GetUploadToken(ctx, req)
	if err != nil {
		logger.Errorf("GetUploadTokenRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}

	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return &api.GetUploadTokenResponse{
		Policy:        resp.Data.Policy,
		Authorization: resp.Data.Authorization,
		Bucket:        resp.Data.Bucket,
		Uid:           req.UserId,
	}, nil
}

func SetUseravatarRPC(ctx context.Context, req *user.SetUserAvatarUrlRequest) error {
	resp, err := userClient.SetUserAvatarUrl(ctx, req)
	if err != nil {
		logger.Errorf("SetUseravatarRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}

	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return nil
}
