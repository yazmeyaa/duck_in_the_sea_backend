package player

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type CreatePlayerDTO struct {
	UserID int
}

type PlayerRepository interface {
	Create(ctx context.Context, data CreatePlayerDTO) (int, error)
	GetByID(ctx context.Context, id int) (*Player, error)
	GetByTelegramID(ctx context.Context, id int64) (*Player, error)
}

type postgresPlayerRepository struct {
	conn *pgx.Conn
}

func (p *postgresPlayerRepository) scanPlayerModel(row pgx.Row) (*Player, error) {
	var player Player
	if err := row.Scan(&player.ID, &player.UserID); err != nil {
		return nil, err
	}

	return &player, nil
}

// GetByTelegramID implements PlayerRepository.
func (p postgresPlayerRepository) GetByTelegramID(ctx context.Context, id int64) (*Player, error) {
	const sql = "SELECT id, user_id FROM player as p LEFT JOIN users as u ON u.id = p.user_id WHERE u.user_id = $1"
	row := p.conn.QueryRow(ctx, sql, id)

	player, err := p.scanPlayerModel(row)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// GetByID implements PlayerRepository.
func (p postgresPlayerRepository) GetByID(ctx context.Context, id int) (*Player, error) {
	const sql = "SELECT id, user_id FROM player WHERE id = $1"
	row := p.conn.QueryRow(ctx, sql, id)

	player, err := p.scanPlayerModel(row)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// Create implements PlayerRepository.
func (p postgresPlayerRepository) Create(ctx context.Context, data CreatePlayerDTO) (int, error) {
	const sql string = "INSERT INTO player (user_id) VALUES ($1) RETURNING id"
	row := p.conn.QueryRow(ctx, sql, data.UserID)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

var _ PlayerRepository = postgresPlayerRepository{}

func NewRepository(conn *pgx.Conn) PlayerRepository {
	return &postgresPlayerRepository{
		conn: conn,
	}
}
