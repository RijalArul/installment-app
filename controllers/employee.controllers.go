package controllers

import (
	"net/http"
	"test-kr-sigma/helpers"
	"test-kr-sigma/models/web"
	"test-kr-sigma/services"

	"github.com/gin-gonic/gin"
)

type EmployeeController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type EmployeeControllerImpl struct {
	employeeService services.EmployeeService
}

func NewEmployeeController(empService services.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{employeeService: empService}
}

func (employeeController *EmployeeControllerImpl) Register(ctx *gin.Context) {
	var inputEmployee web.RegisterEmployeeDTO
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputEmployee)
	} else {
		ctx.ShouldBind(&inputEmployee)
	}

	createEmployee, err := employeeController.employeeService.Register(inputEmployee)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ResponseSuccess(http.StatusCreated, ctx, createEmployee)
}

func (employeeController *EmployeeControllerImpl) Login(ctx *gin.Context) {
	var loginInput web.LoginRequestDTO

	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&loginInput)
	} else {
		ctx.ShouldBind(&loginInput)
	}

	loginEmp, err := employeeController.employeeService.Login(loginInput)
	comparePass := helpers.ComparePass([]byte(loginEmp.Password), []byte(loginInput.Password))
	accessToken := ""

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email & Password Is Invalid"})
		return
	}

	if comparePass == true {
		accessToken = helpers.GenerateToken(loginEmp.ID, loginEmp.Email)
		loginBody := web.UserLoginResponseBody{
			AccessToken: accessToken,
		}
		ResponseSuccess(http.StatusOK, ctx, loginBody)
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email & Password Is Invalid"})
		return
	}

}
