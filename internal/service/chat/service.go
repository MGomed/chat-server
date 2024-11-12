package chat

import (
	"log"

	repository "github.com/MGomed/chat_server/internal/repository"
	db "github.com/MGomed/common/pkg/client/db"
)

type service struct {
	logger    *log.Logger
	repo      repository.Repository
	txManager db.TxManager
}

// NewService is a service constructor
func NewService(logger *log.Logger, repo repository.Repository, txManager db.TxManager) *service {
	return &service{
		logger:    logger,
		repo:      repo,
		txManager: txManager,
	}
}
