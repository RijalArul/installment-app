package services

import (
	"test-kr-sigma/models/entities"
	"test-kr-sigma/models/web"
	"test-kr-sigma/repositories"

	"github.com/gosimple/slug"
)

type GoodService interface {
	Create(goodDTO web.GoodRequestDTO, goodOwnerID uint) (web.GoodResponseBodyDTO, error)
}

type GoodServiceImpl struct {
	goodRepository repositories.GoodRepository
}

func NewGoodService(GoodRepository repositories.GoodRepository) GoodService {
	return &GoodServiceImpl{goodRepository: GoodRepository}
}

func GoodBodyResponse(good entities.Good) web.GoodResponseBodyDTO {
	return web.GoodResponseBodyDTO{
		Name:  good.Name,
		Slug:  good.Slug,
		Price: good.Price,
		Rate:  good.Rate,
	}
}
func (goodService *GoodServiceImpl) Create(goodDTO web.GoodRequestDTO, goodOwnerID uint) (web.GoodResponseBodyDTO, error) {
	good := entities.Good{
		GoodsOwnerID: goodOwnerID,
		Name:         goodDTO.Name,
		Slug:         slug.Make(goodDTO.Name),
		Price:        goodDTO.Price,
		Rate:         goodDTO.Rate,
	}

	createGood, err := goodService.goodRepository.Create(good)
	goodResponseBody := GoodBodyResponse(*createGood)
	return goodResponseBody, err
}
