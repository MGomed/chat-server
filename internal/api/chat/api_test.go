package chat

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	require "github.com/stretchr/testify/require"

	converters "github.com/MGomed/chat_server/internal/converters"
	service_mock "github.com/MGomed/chat_server/internal/service/mocks"
	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

var (
	ctx context.Context

	ctl *gomock.Controller

	mockService *service_mock.MockService

	api *API

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()

	ctl = gomock.NewController(t)
	mockService = service_mock.NewMockService(ctl)

	api = &API{service: mockService}

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
