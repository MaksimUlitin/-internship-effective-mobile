package main

import (
	"effectiveMobileTask/config"
	"effectiveMobileTask/internal/mock"
	"effectiveMobileTask/internal/routes"
	"effectiveMobileTask/internal/storage/database"
	"effectiveMobileTask/lib/logger"
	"log"
)

// @title Music Library API
// @version 1.0
// @description API for managing song information
// @host localhost:8080
// @BasePath /
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	config.LoadConfigEnv()
	logger.Info("environment variables loaded")

	db := database.DbConnect()
	logger.Info("database connect success")

	database.Migrate(db)
	logger.Info("database migrate success")

	go mock.MockServer()
	logger.Info("mock server start success")

	log.Fatal(routes.Router().Run(":" + config.AppConfig.Server.Port))
}
