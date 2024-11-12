package repository

import (
	"context"

	service_model "github.com/MGomed/chat_server/internal/model"
)

//go:generate mockgen -destination=./mocks/repository_mock.go -package=mocks -source=interfaces.go

// Repository declaired interface for database communication
type Repository interface {
	CreateChat(ctx context.Context, name string) (int64, error)
	CreateMembers(ctx context.Context, chatID int64, members []service_model.ChatMember) error
	DeleteChat(ctx context.Context, id int64) error
}
