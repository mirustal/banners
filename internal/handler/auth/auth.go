package auth

import (
	"banners_service/internal/models"
	"banners_service/pkg/config"
	"banners_service/pkg/utils"
	"banners_service/platform/database"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
    var payload models.SignUpInput 

    
    if err := c.BodyParser(&payload); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"description": "bad json",  "error": err.Error()})
    }

    errors := models.ValidateStruct(&payload)
    if len(errors) > 0 { 
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"description": "bad validate json", "errors": errors})
    }

    if payload.Password != payload.PasswordConfirm {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"description": "Passwords do not match"})
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"description": "fail", "message": "Failed to hash the password"})
    }


    newUser := models.User{
        Name:     payload.Name,
        Email:    strings.ToLower(payload.Email),
        Password: string(hashedPassword),
        Role: payload.Role,
    }


    result := database.DB.Create(&newUser)
    if result.Error != nil {
        if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{"description": "fail", "message": "User with that email already exists"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"description": "error", "message": "Something bad happened"})
    }


    config := config.GetConfig()
    now := time.Now().UTC()
    tokenString, err := utils.JwtCreate(newUser)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"description": "fail", "message": "Create JWT"})
    }
    c.Cookie(&fiber.Cookie{
        Name:     "token",
        Value:    tokenString,
        Expires:  now.Add(time.Minute * time.Duration(config.JwtMaxAge)),
        Secure:   false, 
    })

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": tokenString, "description": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
}

func SignInUser(c *fiber.Ctx) error {
    var payload models.SignInInput 


    if err := c.BodyParser(&payload); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"description": "fail", "message": err.Error()})
    }


    if errors := models.ValidateStruct(&payload); len(errors) > 0 { 
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"description": "fail", "errors": errors})
    }

    var user models.User
    result := database.DB.Where("email = ?", strings.ToLower(payload.Email)).First(&user)
    if result.Error != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"description": "fail", "message": "Invalid email or password"})
    }


    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"description": "fail", "message": "Invalid email or password"})
    }

    config := config.GetConfig()
    now := time.Now().UTC()
    tokenString, err := utils.JwtCreate(user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"description": "fail", "message": "Create JWT"})
    }

    c.Cookie(&fiber.Cookie{
        Name:     "token",
        Value:    tokenString,
        Expires:  now.Add(time.Hour * time.Duration(config.JwtMaxAge)),
        Secure:   false, 
    })

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"description": "success", "token": tokenString})
}


func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Millisecond)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"description": "success"})
}



func DeserializeUser(c *fiber.Ctx) error {
    var tokenString string

    authorization := c.Get("Authorization")
    

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else {
		token := c.Get("token")
		if token == "" {
			token = c.Cookies("token")
		}
		tokenString = token
    }

    if tokenString == "" {
        return c.SendStatus(fiber.StatusUnauthorized)
    }


    config := config.GetConfig()


    tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
        if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
        }
        return []byte(config.JwtSecret), nil
    })
    

    if err != nil {
        slog.Error("Invalid token ", err)
        return c.SendStatus(fiber.StatusUnauthorized)
    }
    
    claims, ok := tokenByte.Claims.(jwt.MapClaims)
    if !ok || !tokenByte.Valid {
        slog.Error("Invalid token claim")
        return c.SendStatus(fiber.StatusUnauthorized)
    }

    var user models.User
   
    userID := fmt.Sprintf("%v", claims["sub"]) 
    if result := database.DB.First(&user, "id = ?", userID); result.Error != nil {
        slog.Error("The user belonging to this token no longer exists")
        return c.SendStatus(fiber.StatusUnauthorized)
    }
    c.Locals("user", models.FilterUserRecord(&user))

    return c.Next()
}

func RequireAdminRole(c *fiber.Ctx) error {
    user, ok := c.Locals("user").(models.UserResponse) 
    if !ok || user.Role != "admin" {
        return c.SendStatus(fiber.StatusForbidden)
    }
    return c.Next()
}