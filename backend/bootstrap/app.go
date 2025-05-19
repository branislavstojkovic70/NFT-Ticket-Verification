package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	HttpServer *gin.Engine
	Env        *Env
	Db         *gorm.DB
}

func Run() App {
	app := App{}
	app.Env = NewEnv()
	app.HttpServer = InitHttpServer(app.Env.ServerPort)
	app.Db, _ = InitDB(app.Env.DBHost, app.Env.DBUser, app.Env.DBPass, app.Env.DBName, app.Env.DBPort)
	return app
}
