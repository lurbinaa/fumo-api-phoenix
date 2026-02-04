package main

import (
	"log"
	"net/http"

	"fumo-api/internal/app"
	"fumo-api/internal/db"
)

func main() {
	a := app.NewApplication(app.Config{
		Addr: ":8080",
	})

	a.AddHandler("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("foo bar"))
	})

	log.Printf("At port \"%s\"", a.Config.Addr)

	db.NewDb()

	a.StartListening()

}
