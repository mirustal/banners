package app

import (
	"banners_service/internal/router"
	"banners_service/pkg/config"
	"banners_service/platform/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)


func Init(config *config.Config) {
	database.ConnectDB(config)
	app := fiber.New()
	router.Init(app)
	app.Listen(":8000")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))
}