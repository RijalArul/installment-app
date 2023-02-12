package controllers

import (
	"net/http"
	"test-kr-sigma/helpers"
	"test-kr-sigma/models/web"
	"test-kr-sigma/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GoodController interface {
	Create(ctx *gin.Context)
}

type GoodControllerImpl struct {
	goodService services.GoodService
}

func NewGoodController(GoodService services.GoodService) GoodController {
	return &GoodControllerImpl{goodService: GoodService}
}

func (goodController *GoodControllerImpl) Create(ctx *gin.Context) {
	var goodInput web.GoodRequestDTO
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&goodInput)
	} else {
		ctx.ShouldBind(&goodInput)
	}

	createGood, err := goodController.goodService.Create(goodInput, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ResponseSuccess(http.StatusCreated, ctx, createGood)
}
