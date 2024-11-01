package chat

import (
	"log"

	errors "github.com/MGomed/chat_server/internal/api/errors"
	service "github.com/MGomed/chat_server/internal/service"
	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

// API implements UserAPI grpc server
type API struct {
	chat_api.UnimplementedChatAPIServer

	logger  *log.Logger
	service service.Service
}

// NewAPI is api struct constructor
func NewAPI(logger *log.Logger, service service.Service) *API {
	return &API{
		logger:  logger,
		service: service,
	}
}

func validateName(name string) error {
	if n := len([]rune(name)); n < 2 && n > 32 {
		return errors.ErrNameLenInvalid
	}

	return nil
}
