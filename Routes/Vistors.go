package Routes

import "app/Controllers/Visitors"

func (app RouterApp) visitorsRoutes() {
	app.Gin.GET("/", Visitors.Welcome)
	app.Gin.POST("register", Visitors.Register)
	app.Gin.POST("login", Visitors.Login)

}
