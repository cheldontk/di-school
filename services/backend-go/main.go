package main

import (
	"log"

	"github.com/cheldontk/di-school/di-go/application/database"
	"github.com/cheldontk/di-school/di-go/application/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())
	router.SetupRoutes(app)

	app.Listen(":3000")

}
