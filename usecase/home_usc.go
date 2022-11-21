package usecase

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/woong/conv"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/port"
)

type homeUsc struct {
	beginner        common.TxBeginner
	homeRepo        port.HomeRepo
	shortNoticeRepo port.ShortNoticeRepo
}

func NewHomeUsc(beginner common.TxBeginner, homeRepo port.HomeRepo, shortNoticeRepo port.ShortNoticeRepo) *homeUsc {
	return &homeUsc{
		beginner:        beginner,
		homeRepo:        homeRepo,
		shortNoticeRepo: shortNoticeRepo,
	}
}

func (u *homeUsc) FindByAppConfigID(ctx context.Context, appConfigID string) (dto.Home, error) {
	home, err := u.homeRepo.ReadByAppConfigIDNoTx(ctx, appConfigID)
	if err != nil {
		return dto.NilHome, err
	}

	return conv.ToHomeDto(&home)
}
