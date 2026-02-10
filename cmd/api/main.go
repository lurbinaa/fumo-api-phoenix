package main

import (
	"log"

	"fumo-api/internal/app"
)

func main() {
	a := app.NewApplication(app.Config{
		Addr: ":8080",
	})

	a.RegisterRoutes()
	log.Printf("Listening at port \"%s\"", a.Config.Addr)

	err := a.StartListening()
	if err != nil {
		log.Fatalf("Failed to launch application: %v", err)
	}
}
