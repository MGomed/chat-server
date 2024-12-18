package chat

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	require "github.com/stretchr/testify/require"

	service_model "github.com/MGomed/chat_server/internal/model"
	repo_mocks "github.com/MGomed/chat_server/internal/repository/mocks"
	"github.com/MGomed/common/logger"
)

var (
	ctx context.Context

	ctl *gomock.Controller

	mockRepo *repo_mocks.MockRepository

	serv *service

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()

	ctl = gomock.NewController(t)
	mockRepo = repo_mocks.NewMockRepository(ctl)

	serv = &service{logger: &logger.TestLogger{}, repo: mockRepo}

	t.Cleanup(ctl.Finish)
}

func TestCreate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			chat = service_model.ChatInfo{
				Name: "Chat-1",
				Members: []service_model.ChatMember{
					{
						Name:  "Alex",
						Email: "Alex@mail.com",
					},
					{
						Name:  "Oleg",
						Email: "Oleg@mail.com",
					},
				},
			}

			expectedID int64 = 101
		)

		mockRepo.EXPECT().CreateChat(ctx, chat.Name).Return(expectedID, nil)
		mockRepo.EXPECT().CreateMembers(ctx, expectedID, chat.Members).Return(nil)

		id, err := serv.Create(ctx, &chat)
		require.Equal(t, id, expectedID)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			chat = service_model.ChatInfo{
				Name: "Chat-1",
				Members: []service_model.ChatMember{
					{
						Name:  "Alex",
						Email: "Alex@mail.com",
					},
					{
						Name:  "Oleg",
						Email: "Oleg@mail.com",
					},
				},
			}

			expectedID int64 = 101
		)

		mockRepo.EXPECT().CreateChat(ctx, chat.Name).Return(expectedID, errTest)

		_, err := serv.Create(ctx, &chat)
		require.Equal(t, errTest, err)
	})

	t.Run("Rainy case (couldn't add chat members)", func(t *testing.T) {
		var (
			chat = service_model.ChatInfo{
				Name: "Chat-1",
				Members: []service_model.ChatMember{
					{
						Name:  "Alex",
						Email: "Alex@mail.com",
					},
					{
						Name:  "Oleg",
						Email: "Oleg@mail.com",
					},
				},
			}

			expectedID int64 = 101
		)

		mockRepo.EXPECT().CreateChat(ctx, chat.Name).Return(expectedID, nil)
		mockRepo.EXPECT().CreateMembers(ctx, expectedID, chat.Members).Return(errTest)

		id, err := serv.Create(ctx, &chat)
		require.Equal(t, id, expectedID)
		require.Equal(t, errTest, err)
	})
}

func TestDelete(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockRepo.EXPECT().DeleteChat(ctx, id).Return(nil)

		err := serv.Delete(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockRepo.EXPECT().DeleteChat(ctx, id).Return(errTest)

		err := serv.Delete(ctx, id)
		require.Equal(t, errTest, err)
	})
}
