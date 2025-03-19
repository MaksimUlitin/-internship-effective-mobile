package mock

import (
	"effectiveMobileTask/config"
	"effectiveMobileTask/internal/controllers"
	"effectiveMobileTask/lib/logger"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
)

func MockServer() {
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
