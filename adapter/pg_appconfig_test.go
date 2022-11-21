package adapter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/woong/adapter"
	"github.com/w-woong/woong/entity"
)

var (
	// uat := dnt.Format("20060102150405")
	dnt, _    = time.Parse("20060102150405", "20221120121212")
	appConfig = entity.AppConfig{
		ID:        "b69aa108-f12e-4fa0-bf4f-ba002c11a670",
		Name:      "Test-Woong-App",
		UpdatedAt: &dnt,
	}
)

func TestCreateAppConfig(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPgAppconfig(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	repo.DeleteAppconfig(ctx, tx, appConfig.ID)
	repo.DeleteAppconfig(ctx, tx, appConfig.ID)
	affected, err := repo.CreateAppconfig(ctx, tx, appConfig)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}

func TestReadAppconfig(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	repo := adapter.NewPgAppconfig(gdb)
	res, err := repo.ReadAppconfigNoTx(ctx, appConfig.ID)
	assert.Nil(t, err)
	fmt.Println(res)

}

func TestUpdateAppconfig(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPgAppconfig(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	affected, err := repo.UpdateAppconfig(ctx, tx, appConfig)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}
