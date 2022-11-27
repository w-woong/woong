package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/entity"
	"gorm.io/gorm"
)

type homeGroupProductPg struct {
	db *gorm.DB
}

func NewHomeGroupProductPg(db *gorm.DB) *homeGroupProductPg {
	return &homeGroupProductPg{
		db: db,
	}
}

func (a *homeGroupProductPg) DeleteByHomeID(ctx context.Context, tx common.TxController,
	homeID string) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Where("home_id = ?", homeID).
		Delete(&entity.HomeGroupProduct{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil

}
