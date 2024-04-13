package handler

import (
	"banners_service/internal/models"
	"banners_service/pkg/config"
	"banners_service/platform/database"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestBannerCreate(t *testing.T) {

	testRequest := "/banner"
	app := fiber.New()
	database.ConnectDB(config.GetConfig())
	app.Post(testRequest, BannerCreate)

	trueVal := true
	testBanner := []struct {
		name string
		data models.Banner
		expectedCode int
	}{
		{
			name: "banner create 201",
			data: models.Banner{
				TagIDs:    []int64{1, 2, 3},
				FeatureID: 123,
				Content:    map[string]string{"title": "some_title", "text": "some_text", "url": "some_url"},
				IsActive:  &trueVal,
			},
			expectedCode: 201,
		},
	}

	for _, tt := range testBanner {
		marshalData, err := json.Marshal(tt.data)
		if err != nil {
			t.Fatalf("Failed to marshal userData: %v", err)
		}

		req := httptest.NewRequest("POST", testRequest, bytes.NewBuffer(marshalData))
		resp, err := app.Test(req, -1) 
		if err != nil {
			t.Fatalf("Failed to execute request: %v", err)
		}
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, "TestCase '%s' failed", tt.name)
	}
}