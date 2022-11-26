package port

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/woong/entity"
)

type HomeRepo interface {
	// CreateHome creates home and its associations, but AppConfig.
	CreateHome(ctx context.Context, tx common.TxController, home entity.Home) (int64, error)

	// ReadHomeNoTx reads home and its associations.
	ReadHomeNoTx(ctx context.Context, id string) (home entity.Home, err error)

	// ReadByAppConfigIDNoTx reads home and its associations.
	ReadByAppConfigIDNoTx(ctx context.Context, appConfigID string) (home entity.Home, err error)

	// UpdateHome updates home only.
	UpdateHome(ctx context.Context, tx common.TxController, id string, home entity.Home) (int64, error)

	// DeleteHome deletes home only.
	DeleteHome(ctx context.Context, tx common.TxController, id string) (int64, error)
}

type ShortNoticeRepo interface {
	CreateShortNotice(ctx context.Context, tx common.TxController, notice entity.ShortNotice) (int64, error)
	ReadShortNoticeNoTx(ctx context.Context, id string) (notice entity.ShortNotice, err error)
	ReadByHomeIDNoTx(ctx context.Context, homeID string) (entity.ShortNoticeList, error)
	UpdateShortNotice(ctx context.Context, tx common.TxController, id string, notice entity.Home) (int64, error)
	DeleteShortNotice(ctx context.Context, tx common.TxController, id string) (int64, error)
	DeleteByHomeID(ctx context.Context, tx common.TxController, homeID string) (int64, error)
}

type MainPromotionRepo interface {
	DeleteByHomeID(ctx context.Context, tx common.TxController, homeID string) (int64, error)
}
