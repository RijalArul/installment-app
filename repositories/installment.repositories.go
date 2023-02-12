package repositories

import (
	"test-kr-sigma/models/entities"

	"gorm.io/gorm"
)

type InstallmentRepository interface {
	Create(installment entities.Installment) (*entities.Installment, error)
}

type InstallmentRepositoryImpl struct {
	db *gorm.DB
}

func NewInstallmentRepository(DB *gorm.DB) InstallmentRepository {
	return &InstallmentRepositoryImpl{db: DB}
}

func (i *InstallmentRepositoryImpl) Create(installment entities.Installment) (*entities.Installment, error) {
	err := i.db.Create(&installment).Error
	return &installment, err
}
