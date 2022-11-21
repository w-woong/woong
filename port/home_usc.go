package port

import (
	"context"

	"github.com/w-woong/woong/dto"
)

type HomeUsc interface {
	FindByAppConfigID(ctx context.Context, appConfigID string) (dto.Home, error)
}
