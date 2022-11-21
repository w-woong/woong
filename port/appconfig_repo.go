package port

//go:generate mockgen -destination=./mocks/mock_appconfig_repo.go -package=mocks -mock_names=AppConfigRepo=MockAppConfigRepo -source=./appconfig_repo.go . AppConfigRepo
import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/woong/entity"
)

type AppConfigRepo interface {
	CreateAppconfig(ctx context.Context, tx common.TxController,
		appConfig entity.AppConfig) (int64, error)
	ReadAppconfigNoTx(ctx context.Context, id string) (entity.AppConfig, error)
	UpdateAppconfig(ctx context.Context, tx common.TxController,
		appConfig entity.AppConfig) (int64, error)
	DeleteAppconfig(ctx context.Context, tx common.TxController, id string) (int64, error)
}
