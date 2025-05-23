package bootstrap

import (
	"github.com/branislavstojkovic70/nft-ticket-verification/eth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	HttpServer *gin.Engine
	Env        *Env
	Db         *gorm.DB
	EthClient  *eth.EthClient
}

func Run() App {
	app := App{}
	app.Env = NewEnv()
	app.Db, _ = InitDB(app.Env.DBHost, app.Env.DBUser, app.Env.DBPass, app.Env.DBName, app.Env.DBPort)
	app.HttpServer = InitHttpServer(app.Env.ServerPort, app.Db, app.Env.JWTSecret)
	app.EthClient = eth.Init(app.Env.InfuraUrl, app.Env.ChainID)
	return app
}
