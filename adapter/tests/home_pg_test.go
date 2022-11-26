package adapter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/adapter"
	"github.com/w-woong/woong/entity"
)

var (
	// uat := dnt.Format("20060102150405")
	// dnt, _ = time.ParseInLocation("20060102150405", "20221120121212", time.Local)

	home = entity.Home{
		ID:          "c69aa108-f12e-4fa0-bf4f-ba002c11a671",
		Name:        "Home",
		AppConfigID: appConfig.ID,
		AppConfig:   &appConfig,

		ShortNoticeList: entity.ShortNoticeList{
			entity.ShortNotice{
				ID:          "dd9aa108-f12e-4fa0-bf4f-ba002c11a671",
				HomeID:      "c69aa108-f12e-4fa0-bf4f-ba002c11a671",
				ImgUrl:      "assets/images/icons/2x/outline_notification_important_black_24dp.png",
				Name:        "Delivery",
				Description: "Abount Delivery",
			},
			entity.ShortNotice{
				ID:          "ed9aa108-f12e-4fa0-bf4f-ba002c11a671",
				HomeID:      "c69aa108-f12e-4fa0-bf4f-ba002c11a671",
				ImgUrl:      "assets/images/icons/2x/outline_notification_important_white_24dp.png",
				Name:        "Delivery",
				Description: "Abount Delivery",
			},
		},

		MainPromotionList: entity.MainPromotionList{
			entity.MainPromotion{
				ID:          "dd9aa108-f12e-4fa0-bf4f-ba002c11a671",
				HomeID:      "c69aa108-f12e-4fa0-bf4f-ba002c11a671",
				ImgUrl:      "https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80",
				Name:        "Delivery",
				Description: "Abount Delivery",
				Tags: entity.Tags{
					entity.Tag{
						ID:   "tag2a108-f12e-4fa0-bf4f-ba002c11a671",
						Name: "new",
					},
				},
			},
			entity.MainPromotion{
				ID:          "ed9aa108-f12e-4fa0-bf4f-ba002c11a671",
				HomeID:      "c69aa108-f12e-4fa0-bf4f-ba002c11a671",
				ImgUrl:      "https://images.unsplash.com/photo-1519125323398-675f0ddb6308?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=94a1e718d89ca60a6337a6008341ca50&auto=format&fit=crop&w=1950&q=80",
				Name:        "Delivery",
				Description: "Abount Delivery",
				Tags: entity.Tags{
					entity.Tag{
						ID:   "tag1a108-f12e-4fa0-bf4f-ba002c11a671",
						Name: "old",
					},
				},
			},
		},
	}

	home2 = entity.Home{
		ID:   "c69aa108-0000-4fa0-bf4f-ba002c11a671",
		Name: "Home22",
		// AppConfigID: appConfig.ID,
		// AppConfig:   &appConfig,

		ShortNoticeList: entity.ShortNoticeList{
			entity.ShortNotice{
				ID:          "as9aa108-f12e-1111-bf4f-ba002c11a671",
				HomeID:      "c69aa108-0000-4fa0-bf4f-ba002c11a671",
				ImgUrl:      "assets/images/delivery.png",
				Name:        "Delivery1",
				Description: "Abount Delivery",
			},
			entity.ShortNotice{
				ID:          "as8as108-f12e-1112-bf4f-ba002c11a671",
				HomeID:      "c69aa108-0000-4fa0-bf4f-ba002c11a671",
				ImgUrl:      "assets/images/delivery.png",
				Name:        "Delivery2",
				Description: "Abount Delivery",
			},
		},
	}
)

func Test_homePg_CreateHome(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewHomePg(gdb)
	shortNoticePg := adapter.NewShortNoticePg(gdb)
	mainPromotionPg := adapter.NewMainPromotionPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	mainPromotionPg.DeleteByHomeID(ctx, tx, home.ID)
	shortNoticePg.DeleteByHomeID(ctx, tx, home.ID)
	repo.DeleteHome(ctx, tx, home.ID)
	affected, err := repo.CreateHome(ctx, tx, home)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}

func Test_homePg_CreateHome2(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewHomePg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	affected, err := repo.CreateHome(ctx, tx, home2)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}
func Test_homePg_ReadHomeNoTx(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	repo := adapter.NewHomePg(gdb)
	res, err := repo.ReadHomeNoTx(ctx, home.ID)
	assert.Nil(t, err)
	fmt.Println(res.String())

}

func Test_homePg_ReadByAppConfigIDNoTx(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	repo := adapter.NewHomePg(gdb)
	res, err := repo.ReadByAppConfigIDNoTx(ctx, appConfig.ID)
	assert.Nil(t, err)
	fmt.Println(res.String())

}

func Test_homePg_UpdateHome2(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewHomePg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	affected, err := repo.UpdateHome(ctx, tx, home2.ID, home2)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}
