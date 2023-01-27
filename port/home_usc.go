package port

import (
	"context"

	"github.com/w-woong/woong/dto"
)

type HomeUsc interface {
	AddHome(ctx context.Context, home dto.Home) (int64, error)
	FindByAppConfigID(ctx context.Context, appConfigID string) (dto.Home, error)
}
