package routers

import (
	"net/http"
	"test-kr-sigma/controllers"
	"test-kr-sigma/databases"
	"test-kr-sigma/middlewares"
	"test-kr-sigma/repositories"
	"test-kr-sigma/services"

	"github.com/gin-gonic/gin"
)

//	func uploadFiles(ctx *gin.Context) {
//		fileKTP := middlewares.UploadKTP(ctx)
//		fileSelfie := middlewares.UploadSelfie(ctx)
//		filesRekKoran := middlewares.UploadRekKoran(ctx)
//		ctx.JSON(http.StatusOK, gin.H{
//			"ktp":       fileKTP,
//			"selfie":    fileSelfie,
//			"rek_koran": filesRekKoran,
//		})
//	}

func Routes() *gin.Engine {

	router := gin.Default()
	router.StaticFS("/images", http.Dir("public"))
	getDatabase := databases.GetDB()

	userRepository := repositories.NewUserRepository(getDatabase)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	userRouter := router.Group("/users")

	{
		userRouter.POST("/register", middlewares.ValidateKTP(), middlewares.ValidateSelfie(), userController.Register)
	}
	router.Run()

	return router
}
