package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(server *gin.Engine, db *gorm.DB) {
	RegisterUserRoutes(server, db)
	RegisterEventRoutes(server, db)
	RegisterAdminRoutes(server, db)
	RegisterTicketRoutes(server, db)
	RegisterEventOrganizerRoutes(server, db)
}
