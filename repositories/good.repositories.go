package repositories

import (
	"test-kr-sigma/models/entities"

	"gorm.io/gorm"
)

type GoodRepository interface {
	Create(good entities.Good) (*entities.Good, error)
	FindBySlug(slug string) (*entities.Good, error)
}

type GoodRepositoryImpl struct {
	db *gorm.DB
}

func NewGoodRepository(DB *gorm.DB) GoodRepository {
	return &GoodRepositoryImpl{db: DB}
}

func (g *GoodRepositoryImpl) Create(good entities.Good) (*entities.Good, error) {
	err := g.db.Create(&good).Error
	return &good, err
}

func (g *GoodRepositoryImpl) FindBySlug(slug string) (*entities.Good, error) {
	var good entities.Good
	err := g.db.Model(good).Where("slug = ?", slug).First(&good).Error
	return &good, err
}
