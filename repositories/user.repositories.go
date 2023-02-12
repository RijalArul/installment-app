package repositories

import (
	"test-kr-sigma/databases"
	"test-kr-sigma/models/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	Regsiter(user entities.User, arrRekKoranDTO []*entities.CheckAccount, ctx *gin.Context) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
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
