package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/yazmeyaa/duck_in_the_sea/internal/users/domain"
	"github.com/yazmeyaa/duck_in_the_sea/internal/users/port"
)

type PostgresUserRepository struct {
	conn *pgx.Conn
}

var (
	ErrUserNotFound error = errors.New("user not found")
)

// GetByTelegramID implements port.UserRepository.
func (p PostgresUserRepository) GetByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	const sql string = "SELECT id, name, telegram_id, created_at FROM users WHERE telegram_id = $1"
	row := p.conn.QueryRow(ctx, sql, telegramID)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.TelegramID, &user.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// Create implements port.UserRepository.
func (p PostgresUserRepository) Create(ctx context.Context, data port.CreateUserDTO) (int, error) {
	const sql string = "INSERT INTO users (name, telegram_id) VALUES ($1, $2) RETURNING id"
	row := p.conn.QueryRow(ctx, sql, data.Name, data.TelegramID)
	var id int
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

// GetByID implements port.UserRepository.
func (p PostgresUserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	const sql string = "SELECT id, name, telegram_id, created_at FROM users WHERE id = $1"
	row := p.conn.QueryRow(ctx, sql, id)
	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.TelegramID, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func NewPostgresRepository(pool *pgx.Conn) *PostgresUserRepository {
	return &PostgresUserRepository{
		conn: pool,
	}
}

var _ port.UserRepository = PostgresUserRepository{}
