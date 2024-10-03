package chatapi_usecase

import (
	"context"

	domain "github.com/MGomed/chat_server/internal/domain"
)

type usecase struct {
}

func NewChatAPIUsecase() *usecase {
	return &usecase{}
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
