package Routes

import "app/Controllers/Auth"

func (app RouterApp) authRoutes() {
	app.Gin.GET("/me", Auth.Me)
}
