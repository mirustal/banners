package handler

import (
	"banners_service/internal/handler/auth"
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
	userData :=  models.SignUpInput{ 
		Email:           "test@example@mail.ru",
		Name:            "avito",
		Password:        "123456789",
		PasswordConfirm: "123456789",
	}

	app := fiber.New()
	database.ConnectDB(config.GetConfig())
	app.Post("/auth/register", auth.SignUpUser)
	app.Post("/auth/login", auth.SignInUser)
	app.Post(testRequest, Banner)

	marshalUser, err := json.Marshal(userData)
	if err != nil {
		t.Fatalf("Failed to marshal userData: %v", err)
	}
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(marshalUser))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}
	if resp.StatusCode != fiber.StatusCreated {
		req = httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(marshalUser))
		resp, err = app.Test(req, -1)
		if err != nil {
			t.Fatalf("Failed to execute request: %v", err)
		}
	}

	testBanner := []struct {
		name string
		data models.Banner
		expectedCode int
	}{
		{
			name: "banner create 201",
			data: models.Banner{
				TagIDs:    []int{1, 2, 3},
				FeatureID: 123,
				Content:    map[string]string{"title": "some_title", "text": "some_text", "url": "some_url"},
				IsActive:  true,
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