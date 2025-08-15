package usecase

import (
	"context"
	"log/slog"

	"github.com/yazmeyaa/duck_in_the_sea/internal/users/port"
)

func (u *userUseCase) CreateUser(ctx context.Context, data port.CreateUserDTO) (int, error) {
	id, err := u.repo.Create(ctx, data)
	if err != nil {
		u.logger.ErrorContext(ctx, "Failed to create user", slog.Any("error", err))
		return id, err
	}

	return id, err
}
