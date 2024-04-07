package main

import (
	"banners_service/internal/app"
	"banners_service/pkg/config"
)


func main() {
	// cfgPath := "/Users/mirustal/Documents/project/go/avito_tech/"
	cfg := config.GetConfig()
	app.Init(cfg)
	
}
