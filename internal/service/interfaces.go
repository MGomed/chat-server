package service

import (
	"context"

	service_model "github.com/MGomed/chat_server/internal/model"
)

// Service interface of chat_api usecase
type Service interface {
	Create(ctx context.Context, chat *service_model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, info *service_model.MessageInfo) error
}
