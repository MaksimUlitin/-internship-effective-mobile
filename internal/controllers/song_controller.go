package controllers

import (
	"effectiveMobileTask/internal/models"
	"effectiveMobileTask/internal/storage/database"
	"effectiveMobileTask/lib/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type SongEnriched struct {
	ReleaseDate string `json:"release_date"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type songRequest struct {
	Group string `json:"group" example:"Muse"` // Пример значения
	Song  string `json:"song" example:"Supermassive Black Hole"`
}

// AddSongInfo godoc
// @Summary Add song information
// @Description Add new song information from group and title
// @Tags Songs
// @Accept json
// @Produce json
// @Param request body songRequest true "Request Body"
// @Success 200 {object} models.SongDetail "Song details successfully added"
// @Failure 400 {object} map[string]string "Bad request - missing or invalid parameters"
// @Failure 404 {object} map[string]string "Song not found"
// @Failure 500 {object} map[string]string "Internal server error - database or API error"
// @Router /info [post]
func AddSongInfo(c *gin.Context) {
	var requestBody songRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Error("invalid request body", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	groupName := requestBody.Group
	songTitle := requestBody.Song

	db := database.DbConnect()

	var Group models.Group
	if err := db.Where("name = ?", groupName).FirstOrCreate(&Group, models.Group{Name: groupName}).Error; err != nil {
		logger.Error("failed to find or create artist", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server errors"})
		return
	}

	var song models.Song
	if err := db.Where("group_id = ? AND song = ?", Group.ID, songTitle).First(&song).Error; err != nil {
		logger.Info("song not found", slog.Any("params", map[string]string{"group": groupName, "song": songTitle}))
		songDetail, boolReturn := GetSongDetailAPI(groupName, songTitle, c)
		if boolReturn {
			return
		}

		releaseDate, err := time.Parse("02.01.2006", songDetail.ReleaseDate)
		if err != nil {
			logger.Error("failed to parse release date", slog.Any("error", err))
			releaseDate = time.Now()
		}

		newSong := models.Song{
			GroupId:     Group.ID,
			Title:       songTitle,
			ReleaseDate: releaseDate,
			Text:        songDetail.Text,
			Link:        songDetail.Link,
		}

		if err := db.Create(&newSong).Error; err != nil {
			logger.Error("failed to add new song", slog.Any("error", err), slog.Any("params", map[string]string{"group": groupName, "song": songTitle}))
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			return
		}
		logger.Info("added new song", slog.Any("params", map[string]string{"group": groupName, "song": songTitle}))
		song = newSong
	}

	releaseDateStr := song.ReleaseDate.Format("02.01.2006")

	songDetail := models.SongDetail{
		GroupName:   groupName,
		SongName:    songTitle,
		ReleaseDate: releaseDateStr,
		Text:        song.Text,
		Link:        song.Link,
	}

	SongEnrichFromJSON(&songDetail, groupName, songTitle)
	c.JSON(http.StatusOK, songDetail)
}
