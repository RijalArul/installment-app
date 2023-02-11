package entities

type CheckAccount struct {
	GormModel
	UserID uint
	Name   string `gorm:"not null" valid:"required~Your Rekening Koran is required"`
}
