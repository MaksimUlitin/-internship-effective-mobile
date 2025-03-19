package main

import (
	"effectiveMobileTask/config"
	"effectiveMobileTask/internal/controllers"
	"effectiveMobileTask/internal/routes"
	"effectiveMobileTask/internal/storage/database"
	"effectiveMobileTask/lib/logger"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
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

	go mockServer()
	logger.Info("mock server start success")

	log.Fatal(routes.Router().Run(":" + config.AppConfig.Server.Port))
}

func mockServer() {
	testRouter := gin.Default()

	testRouter.GET("/info", func(c *gin.Context) {
		groupName := c.Query("group")
		songTitle := c.Query("song")

		if groupName == "" || songTitle == "" {
			logger.Debug("group or song is empty")
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
			return
		}

		songDetail, err := controllers.GetSongDetailJSON(groupName, songTitle)
		if err != nil {
			logger.Debug("get song detail fail", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
			return
		}

		logger.Info("request to info succeeded", slog.Any("group", groupName), slog.Any("song", songTitle))
		c.JSON(http.StatusOK, songDetail)
	})

	if err := testRouter.Run(":" + config.AppConfig.Server.MockServerPort); err != nil {
		log.Fatal("error starting mock server", err)
	}
}
