package app

import (
	"net/http"

	"fumo-api/internal/db"
)

type Application struct {
	Config Config
	Db     *db.Database
	mux    *http.ServeMux
}

type Config struct {
	Addr string
}

func NewApplication(cfg Config) *Application {
	db := db.NewDb()
	mux := http.NewServeMux()

	return &Application{
		Config: cfg,
		Db:     db,
		mux:    mux,
	}
}

func (a *Application) AddHandler(route string, f http.HandlerFunc) {
	a.mux.HandleFunc(route, f)
}

func (a *Application) StartListening() error {
	return http.ListenAndServe(a.Config.Addr, a.mux)
}
