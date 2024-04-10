package handler

import (
	"banners_service/internal/models"
	"banners_service/platform/database"


	"github.com/gofiber/fiber/v2"
)
	
func BannerGet(c *fiber.Ctx) error {
    featureID := c.Query("feature_id")
    tagID := c.Query("tag_id")
    limit := c.QueryInt("limit", 10) 
    offset := c.QueryInt("offset", 0)

    query := database.DB.Model(&models.Banner{})

    if featureID != "" {
        query = query.Where("feature_id = ?", featureID)
    }

    if tagID != "" {
        query = query.Where("? = ANY (tag_ids)", tagID) 
    }

    query = query.Limit(limit).Offset(offset)

    var banners []models.Banner
    result := query.Find(&banners)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
    }

	if len(banners) < 1  {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusCreated).JSON(banners)
}

