package handler

import (
	"banners_service/internal/models"
	"banners_service/platform/database"

	"time"

	"github.com/gofiber/fiber/v2"
)

func BannerCreate(c *fiber.Ctx) error {
	var response models.CreateBannerDTO

	if err := c.BodyParser(&response); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	errors := models.ValidateStruct(&response)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}
	var tagIdsInt64 []int64
	for _, val := range response.TagIDs {
		tagIdsInt64 = append(tagIdsInt64, val)
	}

	newBanner := models.Banner{
		TagIDs:    tagIdsInt64,
		FeatureID: response.FeatureID,
		Content:   response.Content,
		IsActive:  response.IsActive,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	result := database.DB.Create(&newBanner)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"banner_id": newBanner.ID})
}
