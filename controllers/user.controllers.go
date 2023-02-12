package controllers

import (
	"fmt"
	"net/http"
	"test-kr-sigma/helpers"
	"test-kr-sigma/middlewares"
	"test-kr-sigma/models/entities"
	"test-kr-sigma/models/web"
	"test-kr-sigma/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	MIMEMultipartPOSTForm = "multipart/form-data"
	appJSON               = "application/json"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
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

	ResponseSuccess(http.StatusCreated, ctx, createUser)
}

func (userController *UserControllerImpl) Login(ctx *gin.Context) {
	var inputLogin web.LoginRequestDTO
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputLogin)
	} else {
		ctx.ShouldBind(&inputLogin)
	}

	loginUser, err := userController.userService.Login(inputLogin)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email & Password Is Invalid"})
		return
	}

	comparePass := helpers.ComparePass([]byte(loginUser.Password), []byte(inputLogin.Password))
	accessToken := ""
	if comparePass == true {
		accessToken = helpers.GenerateToken(loginUser.ID, loginUser.Email, "customer")
		loginBody := web.UserLoginResponseBody{
			AccessToken: accessToken,
		}
		ResponseSuccess(http.StatusOK, ctx, loginBody)

	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email & Password Is Invalid"})
		return
	}

}
func (userController *UserControllerImpl) Update(ctx *gin.Context) {
	var inputExpends web.UpdateExpends
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputExpends)
	} else {
		ctx.ShouldBind(&inputExpends)
	}

	updateExtend, err := userController.userService.Update(inputExpends, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	ResponseSuccess(http.StatusOK, ctx, updateExtend)
}
