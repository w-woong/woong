package port

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/entity"
)

type HomeGroupProductRepo interface {
	CreateHomeGroupProduct(ctx context.Context, tx common.TxController, o entity.HomeGroupProduct) (int64, error)
	DeleteByHomeID(ctx context.Context, tx common.TxController,
		homeID string) (int64, error)
}

type HomeGroupProductUsc interface {
	AddHomeGroupProducts(ctx context.Context, o dto.HomeGroupProductList) (int64, error)
	RemoveByHomeID(ctx context.Context, homeID string) (int64, error)
}
