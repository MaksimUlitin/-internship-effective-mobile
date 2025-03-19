package controllers

import (
	"effectiveMobileTask/internal/models"
	"effectiveMobileTask/internal/storage/database"
	"effectiveMobileTask/lib/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetSongs godoc
// @Summary List songs with optional filtering
// @Description Retrieve a list of songs with optional filtering and pagination
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string false "Filter by Group Name"
// @Param song query string false "Filter by Song Title"
// @Param release_date query string false "Filter by Release Date (format: DD.MM.YYYY)"
// @Param text query string false "Filter by Text"
// @Param link query string false "Filter by Link"
// @Param page query int false "Page number for pagination" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.Song "Songs retrieved successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid parameters"
// @Failure 404 {object} map[string]string "No songs found matching criteria"
// @Failure 500 {object} map[string]string "Internal server error - database error"
// @Router /songs [get]
func GetSongs(c *gin.Context) {
	db := database.DbConnect()
	var songs []models.Song

	group := c.Query("group")
	song := c.Query("song")
	releaseDate := c.Query("release_date")
	text := c.Query("text")
	link := c.Query("link")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}

	limitNumber, err := strconv.Atoi(limit)
	if err != nil || limitNumber < 1 {
		limitNumber = 10
	}

	query := db.Model(&models.Song{}).
		Select("songs.*, groups.name AS group_name").
		Joins("JOIN groups ON songs.group_id = groups.id")

	if group != "" {
		query = query.Where("groups.name ILIKE ?", "%"+group+"%")
	}

	if song != "" {
		query = query.Where("songs.song ILIKE ?", "%"+song+"%")
	}

	if releaseDate != "" {
		date, err := time.Parse("02.01.2006", releaseDate)
		if err == nil {
			query = query.Where("songs.release_date = ?", date)
		}
	}

	if text != "" {
		query = query.Where("songs.text ILIKE ?", "%"+text+"%")
	}

	if link != "" {
		query = query.Where("songs.link ILIKE ?", "%"+link+"%")
	}

	offset := (pageNumber - 1) * limitNumber
	query = query.Offset(offset).Limit(limitNumber)

	if err := query.Find(&songs).Error; err != nil {
		logger.Error("failed to query songs", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to query songs"})
		return
	}

	if len(songs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no songs found matching criteria"})
		return
	}

	logger.Info("Songs retrieved successfully", slog.Int("count", len(songs)))
	c.JSON(http.StatusOK, songs)
}

// GetSongText godoc
// @Summary Get song text by ID with pagination
// @Description Retrieve song text for a specific song ID with pagination support
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param page query int false "Page number for text pagination" default(1)
// @Param limit query int false "Number of text lines per page" default(10)
// @Success 200 {object} map[string]interface{} "Song text retrieved successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid ID format"
// @Failure 404 {object} map[string]string "Song or page not found"
// @Failure 500 {object} map[string]string "Internal server error - database error"
// @Router /songs/{id}/text [get]
func GetSongText(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("invalid song ID format", slog.Any("id", c.Param("id")))
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid song ID format"})
		return
	}

	db := database.DbConnect()
	var song models.Song

	if err := db.Unscoped().First(&song, id).Error; err != nil {
		logger.Error("failed to query song", slog.Any("id", id))
		c.JSON(http.StatusNotFound, gin.H{"message": "song not found"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	text := strings.Split(song.Text, "\n\n")

	totalText := len(text)
	if totalText == 0 {
		logger.Error("no text found for song with id ", slog.Any("id", id))
		c.JSON(http.StatusNotFound, gin.H{"message": "text not found"})
		return
	}

	beginOfIndex := (page - 1) * limit
	endOfIndex := beginOfIndex + limit

	if beginOfIndex >= totalText {
		logger.Error("page out of range for song id", slog.Any("id", id), slog.Any("page", page))
		c.JSON(http.StatusNotFound, gin.H{"message": "no text found for requested page"})
		return
	}

	if endOfIndex > totalText {
		endOfIndex = totalText
	}

	selectText := text[beginOfIndex:endOfIndex]
	resp := map[string]interface{}{
		"songId":    id,
		"page":      page,
		"text":      selectText,
		"total":     totalText,
		"limit":     limit,
		"totalPage": (totalText + limit - 1) / limit,
	}

	logger.Info("retrieved text for song id ", slog.Any("id", id), slog.Any("page", page))
	c.JSON(http.StatusOK, resp)
}
