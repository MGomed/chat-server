package chat

import (
	"context"
)

// Delete removes chat by id
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.repo.DeleteChat(ctx, id)
	if err != nil {
		s.logger.Printf("Failed to delete chat with id - %v from database: %v", id, err)

		return err
	}

	s.logger.Printf("Successfully deleted chat: %v", id)

	return nil
}
