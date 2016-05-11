package application

// Route processing; Need to create controllers
func Route(app *App) {

	dataController := NewDataController(app)

	app.Router.POST("data", dataController.Post)
	app.Router.GET("data/:id", dataController.All)
}
