package usecase

import (
	"context"
	"myreel/app/chat/domain/model"
	"myreel/app/chat/domain/repository"
	"myreel/app/chat/domain/service"
)

type useCase struct {
	db    repository.ChatDB
	svc   service.ChatService
	cache repository.ChatCache
	vRpc  repository.RpcPort
}

type ChatUseCase interface {
	SendMessage(ctx context.Context, senderID, targetID int64, chatType model.ChatType, content string) error
}

func NewChatUseCase(db repository.ChatDB, svc service.ChatService, cache repository.ChatCache, vRpc repository.RpcPort) ChatUseCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
		vRpc:  vRpc,
	}
}
