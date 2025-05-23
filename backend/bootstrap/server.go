package bootstrap

import (
	"github.com/branislavstojkovic70/nft-ticket-verification/api/route"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func InitHttpServer(httpPort string, db *gorm.DB, jwt_secret string) *gin.Engine {
	httpServer := gin.Default()
	route.InitRoutes(httpServer, db, jwt_secret)
	go func() {
		if err := httpServer.Run(":" + httpPort); err != nil {
			panic(err)
		}
	}()
	return httpServer
}
