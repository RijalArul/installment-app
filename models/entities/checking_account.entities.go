package entities

type CheckAccount struct {
	GormModel
	UserID   uint
	RekKoran string `gorm:"not null" valid:"required~Your Rekening Koran is required"`
	User     *User
}
