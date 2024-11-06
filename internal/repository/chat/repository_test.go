package chat

import (
	"context"
	"errors"
	"io"
	"log"
	"testing"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	require "github.com/stretchr/testify/require"

	service_model "github.com/MGomed/chat_server/internal/model"
	repo_errors "github.com/MGomed/chat_server/internal/repository/errors"
	db_mock "github.com/MGomed/chat_server/pkg/client/db/mocks"
)

var (
	ctx    context.Context
	logger *log.Logger

	ctl *gomock.Controller

	mockDB  *db_mock.MockDB
	mockDBC *db_mock.MockClient

	repo *repository

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()
	logger = log.New(io.Discard, "", 0)

	ctl = gomock.NewController(t)
	mockDB = db_mock.NewMockDB(ctl)
	mockDBC = db_mock.NewMockClient(ctl)

	repo = &repository{dbc: mockDBC}

	t.Cleanup(ctl.Finish)
}

func TestCreateChat(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			name = "Chat-1"

			expectedID int64 = 101

			r = &row{
				expectedID: expectedID,
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), gomock.Any()).Return(r)

		id, err := repo.CreateChat(ctx, name)
		require.Equal(t, id, expectedID)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			name = "Chat-1"

			r = &row{
				err: errTest,
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), gomock.Any()).Return(r)

		_, err := repo.CreateChat(ctx, name)
		require.Equal(t, errors.Is(err, repo_errors.ErrQueryExecute), true)
	})
}

func TestCreateMembers(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			chatID int64 = 101

			members = []service_model.ChatMember{
				{
					Name:  "Alex",
					Email: "Alex@mail.com",
				},
				{
					Name:  "Oleg",
					Email: "Oleg@mail.com",
				},
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Query(ctx, gomock.Any(), gomock.Any()).Return(nil, nil)

		err := repo.CreateMembers(ctx, chatID, members)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			chatID int64 = 101

			members = []service_model.ChatMember{
				{
					Name:  "Alex",
					Email: "Alex@mail.com",
				},
				{
					Name:  "Oleg",
					Email: "Oleg@mail.com",
				},
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Query(ctx, gomock.Any(), gomock.Any()).Return(nil, errTest)

		err := repo.CreateMembers(ctx, chatID, members)
		require.Equal(t, errors.Is(err, repo_errors.ErrQueryExecute), true)
	})
}

func TestDeleteChat(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Exec(ctx, gomock.Any(), gomock.Any()).Return(pgconn.CommandTag([]byte("1")), nil)

		err := repo.DeleteChat(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Exec(ctx, gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, errTest)

		err := repo.DeleteChat(ctx, id)
		require.Equal(t, errors.Is(err, repo_errors.ErrQueryExecute), true)
	})
}
