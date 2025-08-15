package usecase

import (
	"context"
	"log/slog"

	"github.com/yazmeyaa/duck_in_the_sea/internal/users/domain"
	"github.com/yazmeyaa/duck_in_the_sea/internal/users/port"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, data port.CreateUserDTO) (int, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	GetUserByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error)
}

type userUseCase struct {
	repo   port.UserRepository
	logger *slog.Logger
}

func NewUserUseCase(repo port.UserRepository, logger *slog.Logger) *userUseCase {
	return &userUseCase{
		repo:   repo,
		logger: logger,
	}
}

var _ UserUseCase = &userUseCase{}
