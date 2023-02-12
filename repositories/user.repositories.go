package repositories

import (
	"test-kr-sigma/databases"
	"test-kr-sigma/models/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Regsiter(user entities.User, arrRekKoranDTO []*entities.CheckAccount, ctx *gin.Context) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(userId uint) (*entities.User, error)
	UpdateExpendAvg(user entities.User, loanLimit entities.LoanLimit) (*entities.LoanLimit, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: Db}
}

func (u *UserRepositoryImpl) Regsiter(user entities.User, arrRekKoranDTO []*entities.CheckAccount, ctx *gin.Context) (*entities.User, error) {
	tx := databases.GetDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {

		return nil, err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var newArrRekKoran []entities.CheckAccount

	for i := 0; i < len(arrRekKoranDTO); i++ {
		newArrRekKoran = append(newArrRekKoran, entities.CheckAccount{
			UserID:   user.ID,
			RekKoran: string(arrRekKoranDTO[i].RekKoran),
		})
	}
	if err := tx.Create(&newArrRekKoran).Error; err != nil {
		tx.Rollback()
		// log.Fatal(err)
	}

	return &user, tx.Commit().Error
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
	var getUser entities.User
	err := u.db.Model(&getUser).Where("email = ?", email).First(&getUser).Error

	return &getUser, err
}

func (u *UserRepositoryImpl) FindByID(userID uint) (*entities.User, error) {
	var getUser entities.User
	err := u.db.Model(&getUser).Where("id = ?", userID).First(&getUser).Error

	return &getUser, err
}

func (u *UserRepositoryImpl) UpdateExpendAvg(user entities.User, loanLimit entities.LoanLimit) (*entities.LoanLimit, error) {
	tx := databases.GetDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := u.db.Preload(clause.Associations).Where("id = ?", user.ID).Updates(user).First(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := u.db.Create(&loanLimit).Error; err != nil {
		tx.Rollback()
	}
	return &loanLimit, tx.Commit().Error
}
