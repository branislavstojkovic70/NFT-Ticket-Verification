package route

import (
	"github.com/branislavstojkovic70/nft-ticket-verification/api/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(server *gin.Engine, db *gorm.DB, jwt_secret string) {
	ctrl := controller.NewAuthController()
	authRoutes := server.Group("/auth")
	{
		authRoutes.POST("/", func(c *gin.Context) {
			ctrl.Login(db, c, jwt_secret)
		})
	}
}
