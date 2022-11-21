package adapter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/adapter"
	"github.com/w-woong/woong/entity"
)

var (
	// uat := dnt.Format("20060102150405")
	// dnt, _ = time.ParseInLocation("20060102150405", "20221120121212", time.Local)
	shortNotice1 = entity.ShortNotice{
		ID:          "109aa108-f12e-1111-bf4f-ba002c11a671",
		HomeID:      "c69aa108-f12e-4fa0-bf4f-ba002c11a671",
		ImgUrl:      "assets/images/delivery.png",
		Name:        "Delivery1",
		Description: "Abount Delivery",
	}

	shortNotice2 = entity.ShortNotice{
		ID:          "208as108-f12e-1112-bf4f-ba002c11a671",
		HomeID:      "c69aa108-f12e-4fa0-bf4f-ba002c11a671",
		ImgUrl:      "assets/images/delivery.png",
		Name:        "Delivery2",
		Description: "Abount Delivery",
	}
)

func Test_shortNoticePg_CreateShortNotice(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewShortNoticePg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	affected, err := repo.CreateShortNotice(ctx, tx, shortNotice1)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}
