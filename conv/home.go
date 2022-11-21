package conv

import (
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.Home{}, &entity.Home{})
	structmapper.StoreMapper(&entity.Home{}, &dto.Home{})

	structmapper.StoreMapper(&dto.ShortNotice{}, &entity.ShortNotice{})
	structmapper.StoreMapper(&entity.ShortNotice{}, &dto.ShortNotice{})
}

func ToHomeEntity(src *dto.Home) (res entity.Home, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToHomeDto(src *entity.Home) (res dto.Home, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToShortNoticeEntity(src *dto.ShortNotice) (res entity.ShortNotice, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToShortNoticeDto(src *entity.ShortNotice) (res dto.ShortNotice, err error) {
	err = structmapper.Map(src, &res)
	return
}
