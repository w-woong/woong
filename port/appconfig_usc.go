package port

//go:generate mockgen -destination=./mocks/mock_appconfig_usc.go -package=mocks -mock_names=AppConfigUsc=MockAppConfigUsc -source=./appconfig_usc.go . AppConfigUsc

import (
	"context"

	"github.com/w-woong/woong/dto"
)

type AppConfigUsc interface {
	AddAppConfig(ctx context.Context, appConfig dto.AppConfig) (int64, error)
	FindAppConfig(ctx context.Context, id string) (dto.AppConfig, error)
	ChangeAppConfig(ctx context.Context, appConfig dto.AppConfig) (int64, error)
}
