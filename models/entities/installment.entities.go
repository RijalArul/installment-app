package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Installment struct {
	GormModel
	OTR               int    `gorm:"not null" valid:"required~Your OTR is required"`
	AdminFee          int    `gorm:"not null" valid:"required~Your Admin Fee is required"`
	AmountInstallment int    `gorm:"not null" valid:"required~Your Installment Amount is required"`
	AmountRate        int    `gorm:"not null" valid:"required~Your Rate Amount is required"`
	GoodName          string `gorm:"not null" valid:"required~Your Goods Name is required"`
	UserID            uint
	GoodID            uint
	LoanLimitID       uint
	User              *User
	Good              *Good
	LoanLimit         *LoanLimit
}

func (i *Installment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(i)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
