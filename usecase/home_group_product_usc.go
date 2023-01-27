package usecase

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/woong/conv"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/port"
)

type homeGroupProductUsc struct {
	beginner common.TxBeginner
	repo     port.HomeGroupProductRepo
}

func NewHomeGroupProductUsc(beginner common.TxBeginner, repo port.HomeGroupProductRepo) *homeGroupProductUsc {
	return &homeGroupProductUsc{
		beginner: beginner,
		repo:     repo,
	}
}

func (u *homeGroupProductUsc) AddHomeGroupProducts(ctx context.Context, o dto.HomeGroupProductList) (int64, error) {
	list, err := conv.ToHomeGroupProductListEntity(o)
	if err != nil {
		return 0, err
	}

	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var affected int64 = 0
	for _, e := range list {
		if e.ID == "" {
			e.CreateSetID()
		}
		res, err := u.repo.CreateHomeGroupProduct(ctx, tx, e)
		if err != nil {
			return 0, err
		}
		affected += res
	}

	return affected, tx.Commit()
}
func (u *homeGroupProductUsc) RemoveByHomeID(ctx context.Context, homeID string) (int64, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	res, err := u.repo.DeleteByHomeID(ctx, tx, homeID)
	if err != nil {
		return 0, err
	}

	return res, tx.Commit()
}
