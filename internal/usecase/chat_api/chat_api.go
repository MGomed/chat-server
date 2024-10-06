package chatapi_usecase

import (
	"context"
	"log"

	domain "github.com/MGomed/chat_server/internal/domain"
)

type usecase struct {
	logger *log.Logger
}

func NewChatAPIUsecase(logger *log.Logger) *usecase {
	return &usecase{
		logger: logger,
	}
}

func (uc *usecase) Create(ctx context.Context, req *domain.CreateRequest) (*domain.CreateResponse, error) {
	// TODO some business logic

	return nil, nil
}

func (uc *usecase) Delete(ctx context.Context, req *domain.DeleteRequest) error {
	// TODO some business logic

	return nil
}

func (uc *usecase) SendMessage(ctx context.Context, req *domain.SendRequest) error {
	// TODO some business logic

	return nil
}
