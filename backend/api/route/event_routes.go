package route

import (
    "github.com/branislavstojkovic70/nft-ticket-verification/api/controller"
    "github.com/branislavstojkovic70/nft-ticket-verification/repository"
    "github.com/branislavstojkovic70/nft-ticket-verification/service"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterEventRoutes(server *gin.Engine, db *gorm.DB) {
    repo := repository.NewEventRepository(db)
    service := service.NewEventService(repo)
    controller := controller.NewEventController(service)

    eventRoutes := server.Group("/event")
    {
        eventRoutes.GET("/", controller.GetAllEvents)
        eventRoutes.GET("/:id", controller.GetEventByID)
        eventRoutes.POST("/", controller.CreateEvent)
        eventRoutes.PUT("/:id", controller.UpdateEvent)
        eventRoutes.DELETE("/:id", controller.DeleteEvent)
    }
}