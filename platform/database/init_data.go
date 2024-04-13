package database

import (
    "banners_service/internal/models"
    "log"
    "golang.org/x/crypto/bcrypt"
)

func InitTestData() {
    users := []models.User{
        {
            Name:     "admin",
            Email:    "admin@example.com",
            Password: hashPassword("123456789"),
            Role:     "admin",
        },
        {
            Name:     "user",
            Email:    "user@example.com",
            Password: hashPassword("123456789"),
            Role:     "user",
        },
    }

    for _, user := range users {
        if err := DB.Create(&user).Error; err != nil {
            log.Printf("Cannot create user %v: %v", user.Email, err)
        }
    }

	firestContent := map[string]string{"test": "first feature"}
	secondContent := map[string]string{"test": "second feature"}

    banners := []models.Banner{
        {
            TagIDs:    []int64{1, 2, 3},
            FeatureID: 1,
            Content:   firestContent,
            IsActive:  pointerToBool(true),
        },
        {
            TagIDs:    []int64{1, 2, 3},
            FeatureID: 2,
            Content:   secondContent,
            IsActive:  pointerToBool(false),
        },
    }

    for _, banner := range banners {
        if err := DB.Create(&banner).Error; err != nil {
            log.Printf("Cannot create banner: %v", err)
        }
    }
}

func hashPassword(password string) string {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatalf("Failed to hash password: %v", err)
    }
    return string(hashedPassword)
}

func pointerToBool(b bool) *bool {
    return &b
}
