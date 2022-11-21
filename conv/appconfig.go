package conv

import (
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.AppConfig{}, &entity.AppConfig{})
	structmapper.StoreMapper(&entity.AppConfig{}, &dto.AppConfig{})
}

func ToAppConfigEntity(src *dto.AppConfig) (res entity.AppConfig, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToAppConfigDto(src *entity.AppConfig) (res dto.AppConfig, err error) {
	err = structmapper.Map(src, &res)
	return
}
