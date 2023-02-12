package entities

import (
	"test-kr-sigma/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Employee struct {
	GormModel
	Email    string `gorm:"not null;unique" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Username string `gorm:"not null;unique" valid:"required~Your username is required"`
	Role     string `gorm:"not null" valid:"required~Your role is required"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(e)

	if errCreate != nil {
		err = errCreate
		return
	}
	e.Password = helpers.HashPass(e.Password)
	err = nil
	return
}
