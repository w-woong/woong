package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/entity"
	"gorm.io/gorm"
)

type PgAppconfig struct {
	db *gorm.DB
}

func NewPgAppconfig(db *gorm.DB) *PgAppconfig {
	return &PgAppconfig{
		db: db,
	}
}

func (a *PgAppconfig) CreateAppconfig(ctx context.Context, tx common.TxController,
	appConfig entity.AppConfig) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&appConfig)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *PgAppconfig) ReadAppconfigNoTx(ctx context.Context, id string) (entity.AppConfig, error) {
	appConfig := entity.AppConfig{}
	res := a.db.WithContext(ctx).
		Where("id = ?", id).
		First(&appConfig)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilAppConfig, txcom.ConvertErr(res.Error)
	}

	return appConfig, nil
}

func (a *PgAppconfig) UpdateAppconfig(ctx context.Context, tx common.TxController,
	appConfig entity.AppConfig) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Model(&appConfig).
		Where("id = ?", appConfig.ID).
		UpdateColumns(entity.AppConfig{UpdatedAt: appConfig.UpdatedAt, Name: appConfig.Name})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgAppconfig) DeleteAppconfig(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.AppConfig{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
