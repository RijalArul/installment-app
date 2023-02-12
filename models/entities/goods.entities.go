package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Good struct {
	GormModel
	Name         string `gorm:"not null;unique" valid:"required~Your Goods Name is required"`
	Slug         string `gorm:"not null" valid:"required~Your Goods Slug Name is required"`
	Price        int    `gorm:"not null" valid:"required~Your Price Goods is required"`
	Rate         int    `gorm:"not null" valid:"required~Your Interest Rate is required"`
	GoodsOwnerID uint
	Installments []Installment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GoodsOwner   *GoodsOwner
}

func (g *Good) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(g)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
