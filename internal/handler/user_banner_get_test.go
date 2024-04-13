package handler

import (
	"banners_service/internal/middleware/cache"
	"banners_service/internal/models"
	"banners_service/pkg/config"
	"banners_service/platform/database"
	"encoding/json"

	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)


func TestUserBannerGet(t *testing.T) {

testRequest := "/user_banner"

app := fiber.New()
database.ConnectDB(config.GetConfig())
app.Get(testRequest, cache.CacheData.Read, UserBannersGet)

testBanner := []struct {
	name string
	query string
	expectedCode int
	expectedFeature int
}{
	{
		name:			"Filter by feature_id = 1",
		query:         "?tag_id=1&feature_id=1",
		expectedCode:  fiber.StatusOK,
		expectedFeature: 1,
	},
	{
		name:          "Filter by feature_id = 2",
		query:         "??tag_id=1&feature_id=2",
		expectedCode:  fiber.StatusOK,
		expectedFeature: 2,
	},

}

for _, tt := range testBanner {
	req := httptest.NewRequest("GET", testRequest+tt.query, nil)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}
	assert.Equalf(t, tt.expectedCode, resp.StatusCode, "TestCase '%s' failed", tt.name)
	if resp.StatusCode == fiber.StatusOK {
		var banner models.Banner
		err := json.NewDecoder(resp.Body).Decode(&banner)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
	

		assert.Equal(t, tt.expectedFeature, banner.FeatureID, banner)
	}
}
}