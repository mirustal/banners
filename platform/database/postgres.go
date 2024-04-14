package database

import (
	"banners_service/internal/models"
	"banners_service/pkg/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUserName, cfg.DBUserPassword, cfg.DBName, cfg.DBPort)

	gormConfig := &gorm.Config{}

	if cfg.LogDB {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	for i := 1; i <= 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
		db, _ := DB.DB()
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}

		log.Printf("Attempt %d failed to connect to db: %v", i, err)
		time.Sleep(time.Duration(i) * time.Second)
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
	}
	err = DB.AutoMigrate(&models.Banner{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
	}

}

func ConnectTestDB(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUserName, cfg.DBUserPassword, cfg.TestDBName, cfg.DBPort)

	gormConfig := &gorm.Config{}
	if cfg.LogDB {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the test db: %s\n", err)
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err = DB.AutoMigrate(&models.User{}, &models.Banner{})
	if err != nil {
		log.Fatal("Migration Failed: ", err)
	}

	InitTestData(DB)
}
