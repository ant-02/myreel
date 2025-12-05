package rpc

import (
	"myreel/app/like/domain/repository"
	"myreel/kitex_gen/video/videoservice"
)

type likeRpcImpl struct {
	video videoservice.Client
}

func NewVideoRpcImpl(v videoservice.Client) repository.RpcPort {
	return &likeRpcImpl{
		video: v,
	}
}

// func (rpc *likeRpcImpl) GetUserIdByUsername(ctx context.Context, username string) (int64, error) {
// 	resp, err := rpc.video.GetUseridByUsername(ctx, &user.GetUserIdByUsernameRequest{
// 		Username: username,
// 	})
// 	if err != nil {
// 		return 0, errno.Errorf(errno.InternalRPCErrorCode, "rpc: failed to get user id by username: %v", err)
// 	}

// 	if resp.Base.Code != errno.SuccessCode {
// 		return 0, errno.NewErrNo(errno.InternalRPCErrorCode, "rpc: failed to get user id by username")
// 	}

// 	return resp.UserId, nil
// }
