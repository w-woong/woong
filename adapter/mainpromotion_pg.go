package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/entity"
	"gorm.io/gorm"
)

type mainPromotionPg struct {
	db *gorm.DB
}

func NewMainPromotionPg(db *gorm.DB) *mainPromotionPg {
	return &mainPromotionPg{
		db: db,
	}
}

func (a *mainPromotionPg) DeleteByHomeID(ctx context.Context, tx common.TxController,
	homeID string) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Where("home_id = ?", homeID).
		Delete(&entity.MainPromotion{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
