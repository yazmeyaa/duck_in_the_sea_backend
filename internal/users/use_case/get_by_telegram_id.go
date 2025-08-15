package usecase

import (
	"context"

	"github.com/yazmeyaa/duck_in_the_sea/internal/users/domain"
)

// GetUserByTelegramID implements UserUseCase.
func (u *userUseCase) GetUserByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	return u.repo.GetByTelegramID(ctx, telegramID)
}
