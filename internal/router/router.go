package router

import (
	"banners_service/internal/handler"
	"banners_service/internal/middleware/auth"
	"banners_service/internal/middleware/cache"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Route("/auth", func(router fiber.Router) {
		router.Post("/register", auth.SignUpUser)
		router.Post("/login", auth.SignInUser)
		router.Get("/logout", auth.DeserializeUser, auth.LogoutUser)
	})
	app.Post("/banner", auth.DeserializeUser, auth.RequireAdminRole, handler.BannerCreate)
	app.Get("/banner", auth.DeserializeUser, auth.RequireAdminRole, handler.BannerGet)
	app.Delete("/banner/:id", auth.DeserializeUser, auth.RequireAdminRole, handler.BannerDel)
	app.Get("/user_banner", auth.DeserializeUser, cache.CacheData.Read, handler.UserBannersGet)
	app.Patch("/banner/:id", auth.DeserializeUser, auth.RequireAdminRole, handler.BannerUpdate)
}
