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

func ToHomeGroupProductListEntity(input dto.HomeGroupProductList) (entity.HomeGroupProductList, error) {
	output := make(entity.HomeGroupProductList, len(input))
	for i := range input {
		c, err := ToHomeGroupProductEntity(&input[i])
		if err != nil {
			return nil, err
		}
		output[i] = c
	}
	return output, nil
}
