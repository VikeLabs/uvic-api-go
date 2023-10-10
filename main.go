package main

import (
	"log"

	_ "github.com/VikeLabs/uvic-api-go/database"
	"github.com/VikeLabs/uvic-api-go/handlers/ssf"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Route("/ssf", ssf.Router)

	addr := "127.0.0.1:8080"
	log.Println("Listening at", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
