package route

import (
	"github.com/branislavstojkovic70/nft-ticket-verification/api/controller"
	"github.com/branislavstojkovic70/nft-ticket-verification/repository"
	"github.com/branislavstojkovic70/nft-ticket-verification/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(server *gin.Engine, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)

	userRoutes := server.Group("/user")
	{
		userRoutes.GET("/", controller.GetAllUsers)
		userRoutes.GET("/:id", controller.GetUserByID)
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
	}
}