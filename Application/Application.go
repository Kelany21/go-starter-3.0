package Application

import (
	"app/Database"
	"app/Database/Seeders"
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) Share() {}

func App() func() *Application {
	return func() *Application {
		var application Application
		application.Gin = gin.Default()
		connectToDataBase(&application)
		err := gotrans.InitLocales("./public/lang")
		if err != nil {
			fmt.Println("Error on load trans files", err.Error())
		}
		err = gotenv.Load(".env")
		if err != nil {
			fmt.Println("Error on load env files", err.Error())
		}
		return &application
	}
}

func NewApp() *Application {
	app := App()
	return app()
}

func (app *Application) Migrate() {
	Database.Migrate(app.DB)
}

func (app *Application) Seed() {
	Seeders.Seeds(app.DB)
}
