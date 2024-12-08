package chat

import (
	repository "github.com/MGomed/chat_server/internal/repository"
	db "github.com/MGomed/common/client/db"
	logger "github.com/MGomed/common/logger"
)

type service struct {
	logger    logger.Interface
	repo      repository.Repository
	txManager db.TxManager
}

// NewService is a service constructor
func NewService(logger logger.Interface, repo repository.Repository, txManager db.TxManager) *service {
	return &service{
		logger:    logger,
		repo:      repo,
		txManager: txManager,
	}
}
