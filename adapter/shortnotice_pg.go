package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/entity"
	"gorm.io/gorm"
)

type shortNoticePg struct {
	db *gorm.DB
}

func NewShortNoticePg(db *gorm.DB) *shortNoticePg {
	return &shortNoticePg{
		db: db,
	}
}

func (a *shortNoticePg) CreateShortNotice(ctx context.Context, tx common.TxController, notice entity.ShortNotice) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&notice)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *shortNoticePg) ReadShortNoticeNoTx(ctx context.Context, id string) (notice entity.ShortNotice, err error) {

	res := a.db.WithContext(ctx).
		Where("id = ?", id).
		Limit(1).Find(&notice)
	if res.Error != nil {
		return entity.NilShortNotice, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.NilShortNotice, common.ErrRecordNotFound
	}
	return
}

func (a *shortNoticePg) ReadByHomeIDNoTx(ctx context.Context, homeID string) (entity.ShortNoticeList, error) {
	list := make(entity.ShortNoticeList, 0, 16)
	res := a.db.WithContext(ctx).
		Where("home_id = ?", homeID).
		Find(&list)
	if res.Error != nil {
		return nil, txcom.ConvertErr(res.Error)
	}
	return list, nil
}

func (a *shortNoticePg) UpdateShortNotice(ctx context.Context, tx common.TxController, id string, notice entity.Home) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		// Session(&gorm.Session{FullSaveAssociations: true}). // update all assossications
		WithContext(ctx).
		Model(&notice).
		Where("id = ?", id).
		Updates(&notice)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil

}

func (a *shortNoticePg) DeleteShortNotice(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Delete(&entity.ShortNotice{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *shortNoticePg) DeleteByHomeID(ctx context.Context, tx common.TxController, homeID string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Where("home_id = ?", homeID).
		Delete(&entity.ShortNotice{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
