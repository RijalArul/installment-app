package services

import (
	"test-kr-sigma/models/entities"
	"test-kr-sigma/models/web"
	"test-kr-sigma/repositories"
)

type GoodOwnerService interface {
	Register(goodOwnerDTO web.GoodsOwnerRegisterDTO) (web.GoodsOwnerResponseBody, error)
	Login(goodOwnerDTO web.LoginRequestDTO) (*entities.GoodsOwner, error)
}

type GoodOwnerServiceImpl struct {
	goodOwnerRepository repositories.GoodOwnerRepository
}

func NewGoodOwnerService(GoodOwnerRepository repositories.GoodOwnerRepository) GoodOwnerService {
	return &GoodOwnerServiceImpl{goodOwnerRepository: GoodOwnerRepository}
}

func ResponseGoodOwnerBody(goodOwner entities.GoodsOwner) web.GoodsOwnerResponseBody {
	return web.GoodsOwnerResponseBody{
		Email: goodOwner.Email,
		Name:  goodOwner.Name,
	}
}

func (goodOwnerService *GoodOwnerServiceImpl) Register(goodOwnerDTO web.GoodsOwnerRegisterDTO) (web.GoodsOwnerResponseBody, error) {
	goodOwner := entities.GoodsOwner{
		Email:    goodOwnerDTO.Email,
		Password: goodOwnerDTO.Password,
		Name:     goodOwnerDTO.Name,
	}
	newGoodOwner, err := goodOwnerService.goodOwnerRepository.Register(goodOwner)
	respBodyGoodOwner := ResponseGoodOwnerBody(*newGoodOwner)
	return respBodyGoodOwner, err
}

func (goodOwnerService *GoodOwnerServiceImpl) Login(goodOwnerDTO web.LoginRequestDTO) (*entities.GoodsOwner, error) {
	loginGoodOwner, err := goodOwnerService.goodOwnerRepository.FindByEmail(goodOwnerDTO.Email)
	return loginGoodOwner, err
}
