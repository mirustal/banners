package handler

import (
	"banners_service/internal/middleware/cache"
	"banners_service/internal/models"
	"banners_service/platform/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserBannersGet(c *fiber.Ctx) error {

	// userToken := c.Get("token")
	// if userToken == "" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Пользователь не авторизован"})
	// }

	tagID := c.QueryInt("tag_id")
	featureID := c.QueryInt("feature_id")

	var banner models.Banner
	query := database.DB.Where("feature_id = ?", featureID).Where("? = ANY(tag_ids)", tagID)
	isActive := true
	user, ok := c.Locals("user").(models.UserResponse)
	if !ok || user.Role != "admin" {
		query.Where("is_Active = ?", &isActive)
	}

	result := query.Find(&banner)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if banner.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	cache.CacheData.Update(strconv.Itoa(tagID)+" "+strconv.Itoa(featureID), banner)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"example": banner.Content,
	})
}
