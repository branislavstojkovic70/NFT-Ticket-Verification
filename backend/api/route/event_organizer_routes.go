package route

import (
    "github.com/branislavstojkovic70/nft-ticket-verification/api/controller"
    "github.com/branislavstojkovic70/nft-ticket-verification/repository"
    "github.com/branislavstojkovic70/nft-ticket-verification/service"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterEventOrganizerRoutes(server *gin.Engine, db *gorm.DB) {
    repo := repository.NewOrganizerRepository(db)
    svc := service.NewOrganizerService(repo)
    ctrl := controller.NewOrganizerController(svc)

    organizerRoutes := server.Group("/event-organizer")
    {
        organizerRoutes.GET("/", ctrl.GetAllOrganizers)
        organizerRoutes.GET("/:id", ctrl.GetOrganizerByID)
        organizerRoutes.POST("/", ctrl.CreateOrganizer)
        organizerRoutes.DELETE("/:id", ctrl.DeleteOrganizer)
    }
}