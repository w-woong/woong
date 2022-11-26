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

	structmapper.StoreMapper(&dto.MainPromotion{}, &entity.MainPromotion{})
	structmapper.StoreMapper(&entity.MainPromotion{}, &dto.MainPromotion{})

	structmapper.StoreMapper(&dto.Tag{}, &entity.Tag{})
	structmapper.StoreMapper(&entity.Tag{}, &dto.Tag{})
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
