package chat

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

// Delete removes chat by id
func (a *API) Delete(ctx context.Context, req *chat_api.DeleteRequest) (*emptypb.Empty, error) {
	err := a.service.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
