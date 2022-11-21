package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/adapter"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/usecase"
)

var (
	// uat := dnt.Format("20060102150405")
	dnt, _    = time.Parse("20060102150405", "20221120121212")
	appConfig = dto.AppConfig{
		ID:        "88ca4470-7690-4073-a0c8-5cc3b84bc2a6",
		Name:      "Test-Woong-App2",
		UpdatedAt: &dnt,
	}
)

func TestAddAppConfigOnline(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPgAppconfig(gdb)
	usc := usecase.NewAppConfigUsc(beginner, repo)

	res, err := usc.AddAppConfig(ctx, appConfig)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, res)
}

func TestChangeAppConfigOnline(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPgAppconfig(gdb)
	usc := usecase.NewAppConfigUsc(beginner, repo)

	res, err := usc.ChangeAppConfig(ctx, appConfig)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, res)

}
