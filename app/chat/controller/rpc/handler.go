package rpc

import (
	"context"
	build "myreel/app/chat/controller/rpc/pack"
	"myreel/app/chat/domain/model"
	"myreel/app/chat/usecase"
	chat "myreel/kitex_gen/chat"
	base "myreel/pkg/base/context"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct {
	useCase usecase.ChatUseCase
}

func NewChatServiceImpl(uc usecase.ChatUseCase) *ChatServiceImpl {
	return &ChatServiceImpl{useCase: uc}
}

// SendMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (resp *chat.SendMessageResponse, err error) {
	resp = new(chat.SendMessageResponse)

	err = s.useCase.SendMessage(ctx, req.SenderId, req.TargetId, model.ChatType(req.ChatType), req.Content)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	return
}

// GetHistory implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetHistory(ctx context.Context, req *chat.GetHistoryRequest) (resp *chat.GetHistoryResponse, err error) {
	resp = new(chat.GetHistoryResponse)

	messages, pagination, err := s.useCase.GetHistoryMessages(ctx, req.UserId, req.TargetId, req.Cursor, req.Limit, model.ChatType(req.ChatType))
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildMessageList(build.BuildMessages(messages), build.BuildPagination(pagination))
	return
}

// GetUnread implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetUnread(ctx context.Context, req *chat.GetUnreadRequest) (resp *chat.GetUnreadResponse, err error) {
	// TODO: Your code here...
	return
}
