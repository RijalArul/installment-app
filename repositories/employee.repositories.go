package repositories

import (
	"test-kr-sigma/models/entities"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Register(employee entities.Employee) (*entities.Employee, error)
	FindByEmail(email string) (*entities.Employee, error)
}

type EmployeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{db: DB}
}

func (e *EmployeeRepositoryImpl) Register(employee entities.Employee) (*entities.Employee, error) {
	err := e.db.Create(&employee).Error
	return &employee, err
}

func (e *EmployeeRepositoryImpl) FindByEmail(email string) (*entities.Employee, error) {
	var employee entities.Employee
	err := e.db.Model(employee).Where("email = ?", email).First(&employee).Error
	return &employee, err
}
