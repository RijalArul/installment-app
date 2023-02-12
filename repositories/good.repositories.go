package repositories

import (
	"test-kr-sigma/models/entities"

	"gorm.io/gorm"
)

type GoodRepository interface {
	Create(good entities.Good) (*entities.Good, error)
}

type GoodRepositoryImpl struct {
	db *gorm.DB
}

func NewGoodRepository(DB *gorm.DB) GoodRepository {
	return &GoodOwnerRepositoryImpl{db: DB}
}

func (g *GoodOwnerRepositoryImpl) Create(good entities.Good) (*entities.Good, error) {
	err := g.db.Create(&good).Error
	return &good, err
}
