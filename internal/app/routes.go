package app

import "fumo-api/internal/handlers"

func (a *Application) RegisterRoutes() {
	a.AddHandler("/test", handlers.TestHandler)
}
