package port

import (
	"context"

	"github.com/yazmeyaa/duck_in_the_sea/internal/game_resources/domain"
)

type CreateItemDTO struct {
	Title       string
	Description string
	Metadata    []any
}

type ResourceManagerRepository interface {
	GetGameItems(ctx context.Context) ([]domain.GameItem[any], error)
	AddGameItem(ctx context.Context, data CreateItemDTO) (domain.GameItem[any], error)
}
