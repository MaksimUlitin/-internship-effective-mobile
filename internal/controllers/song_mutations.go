package controllers

import (
	"effectiveMobileTask/internal/models"
	"effectiveMobileTask/internal/storage/database"
	"effectiveMobileTask/lib/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// UpdateSong godoc
// @Summary Update an existing song
// @Description Update song information by ID (supports partial updates)
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body models.SongUpdate true "Song Update Information (supports partial updates)"
// @Success 200 {object} map[string]string "Song updated successfully"
// @Failure 400 {object} map[string]string "Invalid song data or ID format"
// @Failure 404 {object} map[string]string "Song not found"
// @Failure 500 {object} map[string]string "Internal server error - database error"
// @Router /songs/{id} [patch]
func UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("invalid song ID format", slog.Any("id", c.Param("id")))
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid song ID format"})
		return
	}

	db := database.DbConnect()
	var song models.Song
	if err := db.First(&song, id).Error; err != nil {
		logger.Error("song not found", slog.Any("id", id))
		c.JSON(http.StatusNotFound, gin.H{"message": "song not found"})
		return
	}

	var updateData models.SongUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		logger.Error("invalid song update data", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	updates := make(map[string]interface{})
	updatedFields := make([]string, 0)

	if updateData.GroupName != nil {
		var group models.Group
		if err := db.First(&group, song.GroupId).Error; err != nil {
			logger.Error("group not found", slog.Any("group_id", song.GroupId))
			c.JSON(http.StatusInternalServerError, gin.H{"message": "group not found"})
			return
		}

		if err := db.Model(&group).Update("name", *updateData.GroupName).Error; err != nil {
			logger.Error("failed to update group name", slog.Any("group_id", song.GroupId))
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update group name"})
			return
		}

		updatedFields = append(updatedFields, "group_name")
	}

	if updateData.Song != nil {
		updates["title"] = *updateData.Song
		updatedFields = append(updatedFields, "title")
	}

	if updateData.ReleaseDate != nil {
		// (день.месяц.год)
		date, err := time.Parse("02.01.2006", *updateData.ReleaseDate)
		if err != nil {
			logger.Error("invalid release date format", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid release date format, expected format: DD.MM.YYYY"})
			return
		}
		updates["release_date"] = date
		updatedFields = append(updatedFields, "release_date")
	}

	if updateData.Text != nil {
		updates["text"] = *updateData.Text
		updatedFields = append(updatedFields, "text")
	}

	if updateData.Link != nil {
		updates["link"] = *updateData.Link
		updatedFields = append(updatedFields, "link")
	}

	if len(updates) > 0 {
		if err := db.Model(&song).Updates(updates).Error; err != nil {
			logger.Error("failed to update song", slog.Any("id", id), slog.Any("error", err))
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			return
		}
	}

	if len(updatedFields) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "no updates provided"})
		return
	}

	logger.Info("song and/or group updated successfully", slog.Any("id", id), slog.Any("updated_fields", updatedFields))
	c.JSON(http.StatusOK, gin.H{
		"message":        "song updated successfully",
		"updated_fields": updatedFields,
	})
}

// DeleteSong godoc
// @Summary Delete a song
// @Description Delete a song by its ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} map[string]string "Song deleted successfully"
// @Failure 400 {object} map[string]string "Invalid song ID format"
// @Failure 404 {object} map[string]string "Song not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /songs/{id} [delete]
func DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("invalid song ID format", slog.Any("id", c.Param("id")))
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid song ID format"})
		return
	}

	db := database.DbConnect()
	var song models.Song
	if err := db.First(&song, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Info("song already deleted or does not exist", slog.Any("id", id))
			c.JSON(http.StatusNotFound, gin.H{"message": "song already deleted or does not exist"})
			return
		}
		logger.Error("failed to fetch song", slog.Any("id", id), slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := db.Delete(&song).Error; err != nil {
		logger.Error("failed to delete song", slog.Any("id", id), slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	logger.Info("song deleted successfully", slog.Any("id", id))
	c.JSON(http.StatusOK, gin.H{"message": "song deleted successfully"})
}
