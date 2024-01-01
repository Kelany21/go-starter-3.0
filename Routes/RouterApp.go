package Routes

import "app/Application"

type RouterApp struct {
	*Application.Application
}

func (app RouterApp) Routing() {
	app.visitorsRoutes()
	app.authRoutes()
	app.adminsRoutes()
}
