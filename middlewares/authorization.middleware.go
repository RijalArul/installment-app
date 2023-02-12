package middlewares

import (
	"net/http"
	"test-kr-sigma/databases"
	"test-kr-sigma/models/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminAuthorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		employee := entities.Employee{}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		err := db.Model(employee).Where("id = ?", userID).First(&employee).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You can't access"})
			return
		}

		if employee.Role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You can't access"})
			return
		}
		ctx.Next()
	}
}
