package handler

import (
	"banners_service/internal/models"
	"banners_service/platform/database"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func BannerDel(c *fiber.Ctx) error {
	bannerID := c.Params("id")

	bannerIDInt, err := strconv.Atoi(bannerID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Некорректные данные %v", err),
		})
	}

	result := database.DB.Delete(&models.Banner{}, bannerIDInt)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
