package conv

import (
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.HomeGroupProduct{}, &entity.HomeGroupProduct{})
	structmapper.StoreMapper(&entity.HomeGroupProduct{}, &dto.HomeGroupProduct{})
}

func ToHomeGroupProductEntity(src *dto.HomeGroupProduct) (res entity.HomeGroupProduct, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToHomeGroupProductDto(src *entity.HomeGroupProduct) (res dto.HomeGroupProduct, err error) {
	err = structmapper.Map(src, &res)
	return
}
