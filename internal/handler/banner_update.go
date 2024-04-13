package handler

import (
	"banners_service/internal/models"
	"banners_service/platform/database"
	"strconv"


	"github.com/gofiber/fiber/v2"
)

func BannerUpdate(c *fiber.Ctx) error {

	bannerID := c.Params("id")
	bannerIDInt, err := strconv.Atoi(bannerID)
	if err != nil {

	}
	
    var updateData models.Banner
    if err := c.BodyParser(&updateData); err != nil {
        return c.SendStatus(fiber.StatusBadRequest)
    }
	var origData models.Banner

    query := database.DB.Model(&models.Banner{})
	query = query.Where("id = ?", bannerIDInt)
	result := query.Find(&origData)

	if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
    }

	if origData.ID == 0  {
		return c.SendStatus(fiber.StatusNotFound)
	}
    
	if updateData.FeatureID != 0 {
		origData.FeatureID = updateData.FeatureID
	}
	if updateData.Content != nil {
		origData.Content = updateData.Content
	}
	if len(updateData.TagIDs) > 0 {
		origData.TagIDs = updateData.TagIDs
	}
	if updateData.IsActive != nil {
		origData.IsActive = updateData.IsActive
	}

    result = database.DB.Model(&models.Banner{}).Where("id = ?", bannerIDInt).Updates(origData)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Внутренняя ошибка сервера",
        })
	}
	return c.SendStatus(fiber.StatusOK)
}

func ConvertSlice[E any](in []any) (out []E) {
	for _, v := range in {
        out = append(out, v.(E))
    }
    return
}