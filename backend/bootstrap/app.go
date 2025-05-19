package bootstrap

import (
	"github.com/branislavstojkovic70/nft-ticket-verification/api/route"

	"github.com/gin-gonic/gin"
)

type App struct {
	HttpServer *gin.Engine
}

func Run() App {
	app := App{}
	app.initHttpServer("8000")
	return app
}

func (app *App) initHttpServer(httpPort string) {
	httpServer := gin.Default()
	app.HttpServer = httpServer
	route.InitRoutes(app.HttpServer)
	go func() {
		if err := httpServer.Run(":" + httpPort); err != nil {
			panic(err)
		}
	}()
}