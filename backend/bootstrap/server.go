package bootstrap

import (
	"github.com/branislavstojkovic70/nft-ticket-verification/api/route"

	"github.com/gin-gonic/gin"
)

func InitHttpServer(httpPort string) *gin.Engine {
	httpServer := gin.Default()
	route.InitRoutes(httpServer)
	go func() {
		if err := httpServer.Run(":" + httpPort); err != nil {
			panic(err)
		}
	}()
	return httpServer
}
