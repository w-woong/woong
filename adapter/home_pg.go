package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type homePg struct {
	db *gorm.DB
}

func NewHomePg(db *gorm.DB) *homePg {
	return &homePg{
		db: db,
	}
}

func (a *homePg) CreateHome(ctx context.Context, tx common.TxController, home entity.Home) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&home)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *homePg) ReadHomeNoTx(ctx context.Context, id string) (home entity.Home, err error) {

	res := a.db.WithContext(ctx).
		Where("id = ?", id).
		// Preload("AppConfig").Preload("ShortNoticeList").Preload("MainPromotionList").
		Preload("MainPromotionList.Tags").
		Preload(clause.Associations).
		First(&home)
	if res.Error != nil {
		return entity.NilHome, txcom.ConvertErr(res.Error)
	}

	return
}

func (a *homePg) ReadByAppConfigIDNoTx(ctx context.Context, appConfigID string) (home entity.Home, err error) {

	res := a.db.WithContext(ctx).
		Where("app_config_id = ?", appConfigID).
		// Preload("AppConfig").Preload("ShortNoticeList").Preload("MainPromotionList").Preload("Tags").
		Preload("MainPromotionList.Tags").
		Preload(clause.Associations).
		First(&home)
	if res.Error != nil {
		return entity.NilHome, txcom.ConvertErr(res.Error)
	}

	return
}

func (a *homePg) UpdateHome(ctx context.Context, tx common.TxController, id string, home entity.Home) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		// Session(&gorm.Session{FullSaveAssociations: true}). // update all assossications
		WithContext(ctx).
		Model(&home).
		Where("id = ?", id).
		Updates(&home)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil

}

func (a *homePg) DeleteHome(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Delete(&entity.Home{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
