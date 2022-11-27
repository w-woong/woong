package usecase

import (
	"context"

	"github.com/w-woong/common"
	productdto "github.com/w-woong/product/dto"
	productport "github.com/w-woong/product/port"
	"github.com/w-woong/woong/conv"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/port"
)

type homeUsc struct {
	beginner        common.TxBeginner
	homeRepo        port.HomeRepo
	shortNoticeRepo port.ShortNoticeRepo
	groupSvc        productport.GroupSvc
}

func NewHomeUsc(beginner common.TxBeginner,
	homeRepo port.HomeRepo, shortNoticeRepo port.ShortNoticeRepo,
	groupSvc productport.GroupSvc) *homeUsc {

	return &homeUsc{
		beginner:        beginner,
		homeRepo:        homeRepo,
		shortNoticeRepo: shortNoticeRepo,
		groupSvc:        groupSvc,
	}
}

func (u *homeUsc) FindByAppConfigID(ctx context.Context, appConfigID string) (dto.Home, error) {
	home, err := u.homeRepo.ReadByAppConfigIDNoTx(ctx, appConfigID)
	if err != nil {
		return dto.NilHome, err
	}

	groups := make(productdto.GroupList, 0, len(home.MainProducts))
	for _, mp := range home.MainProducts {
		group, err := u.groupSvc.ReadGroup(ctx, mp.GroupID)
		if err != nil {
			return dto.NilHome, err
		}
		groups = append(groups, group)
	}

	homeDto, err := conv.ToHomeDto(&home)
	if err != nil {
		return dto.NilHome, err
	}

	homeDto.MainProducts = groups
	return homeDto, nil
}
