package application

// Route processing; Need to create controllers
func Route(app *App) {

	dataHistoryController := NewDataHistoryController(app)
	GreenHouseController := NewGreenHouseController(app)

	app.Router.POST("data", dataHistoryController.Post)
	app.Router.POST("green_houses", GreenHouseController.Post)
	app.Router.GET("green_house/:id", GreenHouseController.One)
	app.Router.GET("green_house/:id/history", GreenHouseController.History)
	app.Router.GET("green_house/:id/action", GreenHouseController.Action)
	app.Router.POST("green_house/:id/program", GreenHouseController.Program)
}
