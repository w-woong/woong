package usecase

import (
	"context"
	"time"

	"github.com/w-woong/common"
	"github.com/w-woong/woong/conv"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/port"
)

type appConfigUsc struct {
	beginner common.TxBeginner
	repo     port.AppConfigRepo
}

func NewAppConfigUsc(beginner common.TxBeginner, repo port.AppConfigRepo) *appConfigUsc {
	return &appConfigUsc{
		beginner: beginner,
		repo:     repo,
	}
}

func (u *appConfigUsc) AddAppConfig(ctx context.Context, appConfig dto.AppConfig) (int64, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	appConfigEntity, err := conv.ToAppConfigEntity(&appConfig)
	if err != nil {
		return 0, err
	}
	if appConfigEntity.ID == "" {
		appConfigEntity.CreateSetID()
	}

	rowsAffected, err := u.repo.CreateAppconfig(ctx, tx, appConfigEntity)
	if err != nil {
		return 0, err
	}

	return rowsAffected, tx.Commit()
}

func (u *appConfigUsc) FindAppConfig(ctx context.Context, id string) (dto.AppConfig, error) {
	res, err := u.repo.ReadAppconfigNoTx(ctx, id)
	if err != nil {
		return dto.NilAppConfig, err
	}
	return conv.ToAppConfigDto(&res)
}

func (u *appConfigUsc) ChangeAppConfig(ctx context.Context, appConfig dto.AppConfig) (int64, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	appConfigEntity, err := conv.ToAppConfigEntity(&appConfig)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	appConfigEntity.UpdatedAt = &now

	rowsAffected, err := u.repo.UpdateAppconfig(ctx, tx, appConfigEntity)
	if err != nil {
		return 0, err
	}

	return rowsAffected, tx.Commit()
}
