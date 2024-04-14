package handler

import (
	"banners_service/internal/models"
	"banners_service/pkg/config"
	"banners_service/platform/database"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestBannerGet(t *testing.T) {

	testRequest := "/banner"

	app := fiber.New()
	database.ConnectDB(config.GetConfig())
	app.Post(testRequest, BannerCreate)

	testBanner := []struct {
		name           string
		query          string
		expectedCode   int
		expectedBanner int
	}{
		{
			name:           "Get all",
			query:          "",
			expectedCode:   fiber.StatusOK,
			expectedBanner: 2,
		},
		{
			name:           "Filter by feature_id",
			query:          "?feature_id=1",
			expectedCode:   fiber.StatusOK,
			expectedBanner: 1,
		},
	}

	for _, tt := range testBanner {
		req := httptest.NewRequest("GET", "/banner"+tt.query, nil)

		resp, err := app.Test(req, -1)
		if err != nil {
			t.Fatalf("Failed to execute request: %v", err)
		}
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, "TestCase '%s' failed", tt.name)

		if resp.StatusCode == fiber.StatusOK {
			var banners []models.Banner
			json.NewDecoder(resp.Body).Decode(&banners)
			assert.Len(t, banners, tt.expectedBanner, "TestCase '%s' failed: expected %d banners, got %d", tt.name, tt.expectedBanner, len(banners))
		}
	}
}
