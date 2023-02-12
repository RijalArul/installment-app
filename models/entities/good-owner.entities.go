package entities

import (
	"test-kr-sigma/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type GoodsOwner struct {
	GormModel
	Email    string `gorm:"not null;unique" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Name     string `gorm:"not null;unique" valid:"required~Your company name is required"`
	Good     []Good `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (gow *GoodsOwner) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(gow)

	if errCreate != nil {
		err = errCreate
		return
	}

	gow.Password = helpers.HashPass(gow.Password)

	err = nil
	return
}
