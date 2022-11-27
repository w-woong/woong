package port

import (
	"context"

	"github.com/w-woong/common"
)

type HomeGroupProductRepo interface {
	DeleteByHomeID(ctx context.Context, tx common.TxController,
		homeID string) (int64, error)
}
