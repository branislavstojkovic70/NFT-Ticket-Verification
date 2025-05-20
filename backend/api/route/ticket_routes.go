package route

import (
    "github.com/branislavstojkovic70/nft-ticket-verification/api/controller"
    "github.com/branislavstojkovic70/nft-ticket-verification/repository"
    "github.com/branislavstojkovic70/nft-ticket-verification/service"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterTicketRoutes(server *gin.Engine, db *gorm.DB) {
    repo := repository.NewTicketRepository(db)
    svc := service.NewTicketService(repo)
    ctrl := controller.NewTicketController(svc)

    ticketRoutes := server.Group("/ticket")
    {
        ticketRoutes.GET("/", ctrl.GetAllTickets)
        ticketRoutes.GET("/:id", ctrl.GetTicketByID)
        ticketRoutes.POST("/", ctrl.CreateTicket)
        ticketRoutes.PUT("/:id", ctrl.UpdateTicket)
        ticketRoutes.DELETE("/:id", ctrl.DeleteTicket)
    }
}