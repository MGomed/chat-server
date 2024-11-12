package chat

import (
	"context"
	"errors"
	"io"
	"log"
	"testing"

	gomock "github.com/golang/mock/gomock"
	require "github.com/stretchr/testify/require"

	api_errors "github.com/MGomed/chat_server/internal/api/errors"
	converters "github.com/MGomed/chat_server/internal/converters"
	service_mock "github.com/MGomed/chat_server/internal/service/mocks"
	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

var (
	ctx    context.Context
	logger *log.Logger

	ctl *gomock.Controller

	mockService *service_mock.MockService

	api *API

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()
	logger = log.New(io.Discard, "", 0)

	ctl = gomock.NewController(t)
	mockService = service_mock.NewMockService(ctl)

	api = &API{logger: logger, service: mockService}

	t.Cleanup(ctl.Finish)
}

func TestCreate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			chat = &chat_api.ChatInfo{
				Name: "Chat-1",
			}

			expectedID = int64(101)
		)

		mockService.EXPECT().Create(ctx, converters.ToChatInfoFromAPI(chat)).Return(expectedID, nil)

		resp, err := api.Create(ctx, &chat_api.CreateRequest{Chat: chat})
		require.Equal(t, resp.Id, expectedID)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			chat = &chat_api.ChatInfo{
				Name: "Chat-1",
			}
		)

		mockService.EXPECT().Create(ctx, converters.ToChatInfoFromAPI(chat)).Return(int64(0), errTest)

		_, err := api.Create(ctx, &chat_api.CreateRequest{Chat: chat})
		require.Equal(t, errTest, err)
	})

	t.Run("Rainy case (too short name)", func(t *testing.T) {
		var (
			chat = &chat_api.ChatInfo{
				Name: "C",
			}
		)

		_, err := api.Create(ctx, &chat_api.CreateRequest{Chat: chat})
		require.Equal(t, errors.Is(err, api_errors.ErrNameLenInvalid), true)
	})
}

func TestDelete(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockService.EXPECT().Delete(ctx, id).Return(nil)

		_, err := api.Delete(ctx, &chat_api.DeleteRequest{Id: id})
		require.Equal(t, nil, err)
	})

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockService.EXPECT().Delete(ctx, id).Return(errTest)

		_, err := api.Delete(ctx, &chat_api.DeleteRequest{Id: id})
		require.Equal(t, errTest, err)
	})
}
