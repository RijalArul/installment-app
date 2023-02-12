package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type CheckAccount struct {
	GormModel
	UserID   uint
	RekKoran string `gorm:"not null" valid:"required~Your Rekening Koran is required"`
	User     *User
}

func (ca *CheckAccount) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(ca)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
