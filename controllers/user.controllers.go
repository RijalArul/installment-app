package controllers

import (
	"fmt"
	"net/http"
	"test-kr-sigma/helpers"
	"test-kr-sigma/middlewares"
	"test-kr-sigma/models/entities"
	"test-kr-sigma/models/web"
	"test-kr-sigma/services"

	"github.com/gin-gonic/gin"
)

const (
	MIMEMultipartPOSTForm = "multipart/form-data"
)

type UserController interface {
	Register(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func ResponseSuccess(code int, ctx *gin.Context, data interface{}) {
	ctx.JSON(code, gin.H{
		"data": data,
	})
}

func NewUserController(UserService services.UserService) UserController {
	return &UserControllerImpl{userService: UserService}
}

func (userController *UserControllerImpl) Register(ctx *gin.Context) {
	var inputUser web.UserRegisterDTO
	var inputCheck web.CheckAccountDTO

	contentType := helpers.GetContentType(ctx)

	if contentType == MIMEMultipartPOSTForm {
		ctx.ShouldBind(&inputUser)
	} else {
		ctx.ShouldBind(&inputUser)
	}

	if contentType == MIMEMultipartPOSTForm {
		ctx.ShouldBind(&inputCheck)
	} else {
		ctx.ShouldBind(&inputCheck)
	}

	if err := ctx.ShouldBind(&inputUser); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBind(&inputCheck); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// userEntity.KTP = middlewares.UploadKTP(ctx)
	// userEntity.Selfie = middlewares.UploadSelfie(ctx)

	rekKoran := middlewares.UploadRekKoran(ctx)
	arrRekKoran := []*entities.CheckAccount{}
	if len(rekKoran) > 0 {
		for i := 0; i < len(rekKoran); i++ {
			arrRekKoran = append(arrRekKoran, &entities.CheckAccount{
				RekKoran: string(rekKoran[i]),
			})
		}
	}
	createUser, err := userController.userService.Register(inputUser, arrRekKoran, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if err

	ResponseSuccess(http.StatusCreated, ctx, createUser)
}
