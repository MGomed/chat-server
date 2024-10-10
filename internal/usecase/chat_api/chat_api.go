package chatapi_usecase

import (
	"context"
	"log"

	domain "github.com/MGomed/chat_server/internal/domain"
)

type usecase struct {
	logger *log.Logger
}

// NewChatAPIUsecase is a usecase constructor
func NewChatAPIUsecase(logger *log.Logger) *usecase {
	return &usecase{
		logger: logger,
	}
}

// Create creates new chat
func (uc *usecase) Create(_ context.Context, _ *domain.CreateRequest) (*domain.CreateResponse, error) {
	// TODO some business logic

	return nil, nil
}

// Delete removes chat by id
func (uc *usecase) Delete(_ context.Context, _ *domain.DeleteRequest) error {
	// TODO some business logic

	return nil
}

// SendMessage send user's message to chat
func (uc *usecase) SendMessage(_ context.Context, _ *domain.SendRequest) error {
	// TODO some business logic

	return nil
}
