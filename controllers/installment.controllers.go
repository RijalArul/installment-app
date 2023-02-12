package controllers

import (
	"net/http"
	"strconv"
	"test-kr-sigma/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type InstallmentController interface {
	Create(ctx *gin.Context)
}

type InstallmentControllerImpl struct {
	installmentService services.InstallmentService
}

func NewInstallmentController(InstallmentService services.InstallmentService) InstallmentController {
	return &InstallmentControllerImpl{installmentService: InstallmentService}
}

func (installmentController *InstallmentControllerImpl) Create(ctx *gin.Context) {
	loanLimitID := ctx.Param("loanLimitID")
	goodSlug := ctx.Param("goodSlug")
	selectMonth := ctx.Query("select_month")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	loanLimitParse, _ := strconv.ParseUint(loanLimitID, 10, 32)
	userID := uint(userData["id"].(float64))

	createInstallment, err := installmentController.installmentService.Create(uint(loanLimitParse), goodSlug, userID, selectMonth)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Your Loan / Good Not Found"})
		return
	}

	ResponseSuccess(http.StatusCreated, ctx, createInstallment)
}
