package chatapi_usecase

import (
	"context"
	"log"

	domain "github.com/MGomed/chat_server/internal/domain"
)

// DatabaseAdapter declaired interface for database communication
type DatabaseAdapter interface {
	CreateChat(ctx context.Context, info *domain.ChatInfo) (int, error)
	DeleteChat(ctx context.Context, id int) (int, error)
}

type usecase struct {
	logger    *log.Logger
	dbAdapter DatabaseAdapter
}

// NewChatAPIUsecase is a usecase constructor
func NewChatAPIUsecase(logger *log.Logger, dbAdapter DatabaseAdapter) *usecase {
	return &usecase{
		logger:    logger,
		dbAdapter: dbAdapter,
	}
}

// Create creates new chat
func (uc *usecase) Create(ctx context.Context, req *domain.CreateRequest) (*domain.CreateResponse, error) {
	id, err := uc.dbAdapter.CreateChat(ctx, req.Info)
	if err != nil {
		uc.logger.Printf("Failed to add chat %v in database: %v", req.Info, err)

		return nil, err
	}

	uc.logger.Printf("Successfully added chat with id: %v", id)

	return &domain.CreateResponse{
		ID: int64(id),
	}, nil
}

// Delete removes chat by id
func (uc *usecase) Delete(ctx context.Context, req *domain.DeleteRequest) error {
	_, err := uc.dbAdapter.DeleteChat(ctx, int(req.ID))
	if err != nil {
		uc.logger.Printf("Failed to delete chat with id - %v from database: %v", req.ID, err)

		return err
	}

	uc.logger.Printf("Successfully deleted chat: %v", req.ID)

	return nil
}

// SendMessage send user's message to chat
func (uc *usecase) SendMessage(_ context.Context, _ *domain.SendRequest) error {
	// TODO some business logic

	return nil
}
