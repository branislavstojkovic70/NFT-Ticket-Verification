package bootstrap

import (
	"time"

	"github.com/branislavstojkovic70/nft-ticket-verification/api/route"
	"github.com/gin-contrib/cors"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func InitHttpServer(httpPort string, db *gorm.DB, jwt_secret string) *gin.Engine {
	httpServer := gin.Default()
	httpServer.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	route.InitRoutes(httpServer, db, jwt_secret)
	go func() {
		if err := httpServer.Run(":" + httpPort); err != nil {
			panic(err)
		}
	}()
	return httpServer
}
