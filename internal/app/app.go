package app

import (
	"banners_service/internal/router"
	"banners_service/pkg/config"
	"banners_service/pkg/logger"
	"banners_service/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func Init(config *config.Config, logger *logger.Logger) {
	database.ConnectDB(config)
	app := fiber.New()
	router.Init(app)
	app.Listen(":8000")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))
}