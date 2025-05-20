package route

import (
    "github.com/branislavstojkovic70/nft-ticket-verification/api/controller"
    "github.com/branislavstojkovic70/nft-ticket-verification/repository"
    "github.com/branislavstojkovic70/nft-ticket-verification/service"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterAdminRoutes(server *gin.Engine, db *gorm.DB) {
    repo := repository.NewAdminRepository(db)
    svc := service.NewAdminService(repo)
    ctrl := controller.NewAdminController(svc)

    adminRoutes := server.Group("/admin")
    {
        adminRoutes.GET("/", ctrl.GetAllAdmins)
        adminRoutes.GET("/:id", ctrl.GetAdminByID)
        adminRoutes.POST("/", ctrl.CreateAdmin)
        adminRoutes.PUT("/:id", ctrl.UpdateAdmin)
        adminRoutes.DELETE("/:id", ctrl.DeleteAdmin)
    }
}