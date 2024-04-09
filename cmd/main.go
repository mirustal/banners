package main

import (
	"banners_service/internal/app"
	"banners_service/pkg/config"
	"banners_service/pkg/logger"
)


func main() {
	// cfgPath := "/Users/mirustal/Documents/project/go/avito_tech/"
	cfg := config.GetConfig()
	log := logger.SetupLogger("debug")
	app.Init(cfg, log)
	
}
