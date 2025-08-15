package usecase

import (
	"context"

	"github.com/yazmeyaa/duck_in_the_sea/internal/users/domain"
)

func (u *userUseCase) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}
