package handler

import (
	"banners_service/internal/models"
	"banners_service/platform/database"

	"github.com/gofiber/fiber/v2"
)



func UserBannersGet(c *fiber.Ctx) error {

	userToken := c.Get("token")
	if userToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Пользователь не авторизован"})
	}

	tagID := c.QueryInt("tag_id")
	featureID := c.QueryInt("feature_id")
	useLastRevision := c.QueryBool("use_last_revision")

	var banner models.Banner
	query := database.DB.Where("feature_id = ?", featureID).Where("? = ANY(tag_ids)", tagID)

	if useLastRevision {
		// Применение логики для получения актуальной информации, если требуется
	}

	result := query.Find(&banner)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
    }

	if banner.ID == 0  {
		return c.SendStatus(fiber.StatusNotFound)
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"example": "123",
			"queries": banner,
	})
}