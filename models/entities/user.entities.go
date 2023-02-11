package entities

import (
	"test-kr-sigma/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	NIK           string `gorm:"not null;unique" valid:"required~Your NIK is required,minstringlength(16)~NIK has to have a minimum length of 16 characters"`
	Email         string `gorm:"not null;unique" valid:"required~Your email is required,email~Invalid email format"`
	Password      string `gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Salary        int    `gorm:"not null" valid:"required~Your Gaji is required"`
	BirthPlace    string `gorm:"not null" valid:"required~Your Tempat Lahir is required"`
	BirthDate     string `gorm:"not null" valid:"required~Your Tanggal Lahir is required"`
	KTP           string `gorm:"not null" valid:"required~Your KTP is required"`
	Selfie        string `gorm:"not null" valid:"required~Your Selfie is required"`
	ExpendAverage int    `gorm:"not null"`
	CheckAccounts []CheckAccount
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
