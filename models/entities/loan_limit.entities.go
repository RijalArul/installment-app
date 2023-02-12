package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type LoanLimit struct {
	GormModel
	UserID       uint
	FirstMonth   int           `gorm:"not null" valid:"required~Your First Tenor is required"`
	SecondMonth  int           `gorm:"not null" valid:"required~Your Second Tenor is required"`
	ThirdMonth   int           `gorm:"not null" valid:"required~Your Third Tenor is required"`
	FourthMonth  int           `gorm:"not null" valid:"required~Your Fourth Tenor is required"`
	Installments []Installment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User         *User
}

func (ll *LoanLimit) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(ll)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
