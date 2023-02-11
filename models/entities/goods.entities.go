package entities

type Good struct {
	GormModel
	Name  string `gorm:"not null;unique" valid:"required~Your Goods Name is required"`
	Slug  string `gorm:"not null" valid:"required~Your Goods Slug Name is required"`
	Price int    `gorm:"not null" valid:"required~Your Price Goods is required"`
	Rate  int    `gorm:"not null" valid:"required~Your Interest Rate is required"`
}
