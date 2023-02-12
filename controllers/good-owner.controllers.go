package controllers

import (
	"net/http"
	"test-kr-sigma/helpers"
	"test-kr-sigma/models/web"
	"test-kr-sigma/services"

	"github.com/gin-gonic/gin"
)

type GoodsOwnerController interface {
	Register(ctx *gin.Context)
}

type GoodsOwnerControllerImpl struct {
	goodsOwnerService services.GoodOwnerService
}

func NewGoodOwnerController(GoodOwnerService services.GoodOwnerService) GoodsOwnerController {
	return &GoodsOwnerControllerImpl{goodsOwnerService: GoodOwnerService}
}

func (goodsOwnerController *GoodsOwnerControllerImpl) Register(ctx *gin.Context) {
	var goodOwnerInput web.GoodsOwnerRegisterDTO
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&goodOwnerInput)
	} else {
		ctx.ShouldBind(&goodOwnerInput)
	}

	createGoodsOwner, err := goodsOwnerController.goodsOwnerService.Register(goodOwnerInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ResponseSuccess(http.StatusCreated, ctx, createGoodsOwner)
}
