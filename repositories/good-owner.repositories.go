package repositories

import (
	"test-kr-sigma/models/entities"

	"gorm.io/gorm"
)

type GoodOwnerRepository interface {
	Register(goodOwner entities.GoodsOwner) (*entities.GoodsOwner, error)
}

type GoodOwnerRepositoryImpl struct {
	db *gorm.DB
}

func NewGoodOwnerRepository(DB *gorm.DB) GoodOwnerRepository {
	return &GoodOwnerRepositoryImpl{db: DB}
}

func (gow *GoodOwnerRepositoryImpl) Register(goodOwner entities.GoodsOwner) (*entities.GoodsOwner, error) {
	err := gow.db.Create(&goodOwner).Error
	return &goodOwner, err
}
