package auth

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



func TestAuth(t *testing.T) {

	type args struct {
		userData    models.SignUpInput
		route       string
	}
	Email := "1@gmail.ru"
	testsRegister := []struct {
		name         string
		args         args
		expectedCode int
	}{
		{
			name: "register user 201", 
			args: args{
				userData: models.SignUpInput{ 
					Email:           Email,
					Name:            "avito",
					Password:        "123456789",
					PasswordConfirm: "123456789",
				},
				route: "/auth/register", 
			},
			expectedCode: 201, 
		},
		{
			name: "duplicate email 409",
			args: args{
				userData: models.SignUpInput{ 
					Email:           Email,
					Name:            "avito",
					Password:        "123456789",
					PasswordConfirm: "123456789",
				},
				route: "/auth/register", 
			},
			expectedCode: 409,
		},
	}

	testLogout := []struct {
		name string
		route string
		expectedCode int
	}{
		{
			name: "logout", 
			route: "/auth/logout",
			expectedCode: 200, 
		},
	}

	testsLogin := []struct {
		name         string
		args         args
		expectedCode int
	}{
		{
			name: "login user 200", 
			args: args{
				userData: models.SignUpInput{ 
					Email:           Email,
					Password:        "123456789",
				},
				route: "/auth/login", 
			},
			expectedCode: 200, 
		},
		{
			name: "Fail Login 401",
			args: args{
				userData: models.SignUpInput{ 
					Email:           "empty",
					Name:            "avito",
					Password:        "123456789",
					PasswordConfirm: "123456789",
				},
				route: "/auth/login", 
			},
			expectedCode: 401,
		},
	}
	

	app := fiber.New()
	database.ConnectDB(config.GetConfig())
	app.Post("/auth/register", SignUpUser)
	app.Post("/auth/login", SignInUser)
	app.Get("/auth/logout", LogoutUser)


	for _, tt := range testsRegister {
		userData, err := json.Marshal(tt.args.userData)
		if err != nil {
			t.Fatalf("Failed to marshal userData: %v", err)
		}

		req := httptest.NewRequest("POST", tt.args.route, bytes.NewBuffer(userData))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1) 
		if err != nil {
			t.Fatalf("Failed to execute request: %v", err)
		}
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, "TestCase '%s' failed", tt.name)
	}

	for _, tt := range testLogout {
		
		req := httptest.NewRequest("GET", tt.route, nil)
		resp, err := app.Test(req, -1) 
		if err != nil {
			t.Fatalf("Failed to execute request: %v", err)
		}
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, "TestCase '%s' failed", tt.name)
	}

	for _, tt := range testsLogin {
		userData, err := json.Marshal(tt.args.userData)
		if err != nil {
			t.Fatalf("Failed to marshal userData: %v", err)
		}

		req := httptest.NewRequest("POST", tt.args.route, bytes.NewBuffer(userData))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1) 
		if err != nil {
			t.Fatalf("Failed to execute request: %v", err)
		}
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, "TestCase '%s' failed", tt.name)
	}
}



