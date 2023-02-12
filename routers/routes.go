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
		userRouter.POST("/login", userController.Login)
		userRouter.Use(middlewares.Authenthication())
		userRouter.Use(middlewares.AdminAuthorize())
		userRouter.PUT("/:id", userController.Update)
	}

	employeeRepository := repositories.NewEmployeeRepository(getDatabase)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)
	employeeRouter := router.Group("/employees")
	{
		employeeRouter.POST("/register", employeeController.Register)

		employeeRouter.POST("/login", employeeController.Login)
	}

	goodsOwnerRepository := repositories.NewGoodOwnerRepository(getDatabase)
	goodsOwnerService := services.NewGoodOwnerService(goodsOwnerRepository)
	goodsOwnerController := controllers.NewGoodOwnerController(goodsOwnerService)
	goodsOwnerRouter := router.Group("/goodsOwners")

	{
		goodsOwnerRouter.POST("/register", goodsOwnerController.Register)
	}
	router.Run()

	return router
}
