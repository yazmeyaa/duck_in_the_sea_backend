package port

import (
	"context"

	"github.com/yazmeyaa/duck_in_the_sea/internal/users/domain"
)

type CreateUserDTO struct {
	Name       string
	TelegramID int
}

type UserRepository interface {
	Create(ctx context.Context, data CreateUserDTO) (int, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error)
}
