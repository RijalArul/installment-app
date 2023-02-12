package services

import (
	"test-kr-sigma/middlewares"
	"test-kr-sigma/models/entities"
	"test-kr-sigma/models/web"
	"test-kr-sigma/repositories"

	"github.com/gin-gonic/gin"
)

func ResponseBodyUser(user *entities.User) web.UserResponseBodyDTO {
	return web.UserResponseBodyDTO{
		NIK:        user.NIK,
		Email:      user.Email,
		Salary:     user.Salary,
		BirthPlace: user.BirthPlace,
		BirthDate:  user.BirthDate,
		KTP:        user.KTP,
		Selfie:     user.Selfie,
	}
}

func ResponseBodyLogin(accessToken string) web.UserLoginResponseBody {
	return web.UserLoginResponseBody{
		AccessToken: accessToken,
	}
}

type UserService interface {
	Register(userDto web.UserRegisterDTO, arrCheckKoran []*entities.CheckAccount, ctx *gin.Context) (web.UserResponseBodyDTO, error)
	Login(userDTO web.UserLoginRequestDTO) (*entities.User, error)
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepo}
}

func (userService *UserServiceImpl) Register(userDto web.UserRegisterDTO, arrCheckKoran []*entities.CheckAccount, ctx *gin.Context) (web.UserResponseBodyDTO, error) {
	user := entities.User{
		NIK:           userDto.NIK,
		Email:         userDto.Email,
		Password:      userDto.Password,
		Salary:        userDto.Salary,
		BirthPlace:    userDto.BirthPlace,
		BirthDate:     userDto.BirthDate,
		KTP:           middlewares.UploadKTP(ctx),
		Selfie:        middlewares.UploadSelfie(ctx),
		ExpendAverage: 0,
	}

	createUser, err := userService.userRepository.Regsiter(user, arrCheckKoran, ctx)

	userBody := web.UserResponseBodyDTO{
		NIK:        createUser.NIK,
		Email:      createUser.Email,
		Salary:     createUser.Salary,
		BirthPlace: createUser.BirthPlace,
		BirthDate:  createUser.BirthDate,
		KTP:        createUser.KTP,
		Selfie:     createUser.Selfie,
	}
	return userBody, err
}

func (userService *UserServiceImpl) Login(userDTO web.UserLoginRequestDTO) (*entities.User, error) {
	login, err := userService.userRepository.FindByEmail(userDTO.Email)

	return login, err
}
