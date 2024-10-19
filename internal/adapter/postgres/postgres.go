package postgres

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/MGomed/chat_server/internal/domain"
	sq "github.com/Masterminds/squirrel"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

const (
	chatTable       = "chat"
	chatMemberTable = "chat_member"
)

var (
	errQueryBuild   = errors.New("failed to build query")
	errQueryExecute = errors.New("failed to execute query")
)

// Config declair postgres config interface
type Config interface {
	DSN() string
}

type adapter struct {
	pool *pgxpool.Pool
}

// NewAdapter is adapter struct constructor
func NewAdapter(ctx context.Context, config Config) (*adapter, error) {
	pool, err := pgxpool.Connect(ctx, config.DSN())
	if err != nil {
		return nil, err
	}

	return &adapter{
		pool: pool,
	}, nil
}

func (a *adapter) CreateChat(ctx context.Context, info *domain.ChatInfo) (int, error) {
	builder := sq.Insert(chatTable).
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(info.Name).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	var chatID int
	err = a.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	if len(info.Members) != 0 {
		if err := a.createChatMembers(ctx, chatID, info.Members); err != nil {
			return 0, err
		}
	}

	return chatID, nil
}

func (a *adapter) createChatMembers(ctx context.Context, chatID int, members []domain.ChatMember) error {
	builder := sq.Insert(chatMemberTable).
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "name", "email")

	for _, member := range members {
		builder = builder.Values(chatID, member.Name, member.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	_, err = a.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return nil
}

func (a *adapter) DeleteChat(ctx context.Context, id int) (int, error) {
	builder := sq.Delete(chatTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	res, err := a.pool.Exec(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return int(res.RowsAffected()), nil
}
