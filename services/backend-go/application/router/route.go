package router

import (
	"github.com/cheldontk/di-school/di-go/application/handler"
	"github.com/cheldontk/di-school/di-go/application/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/api",
		},
		StatusCode: 301,
	}))
	//Middleware
	api := app.Group("/api", middleware.AuthReq())

	//routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}
