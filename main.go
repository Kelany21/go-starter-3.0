package main

import (
	"app/Application"
	"app/Routes"
)

/// gin framework (routing , validate requests , response  , middleware)
/// gorm (connect database , quiers)
/// auth user
/// lanuages
func main() {
	app := Application.NewApp()
	/// migrate project
	app.Migrate()
	/// seed data
	app.Seed()
	/// close app Connection
	Application.CloseConnection(app)
	/// start routeing
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()
	/// start server app
	app.Gin.Run(":9090")
}
