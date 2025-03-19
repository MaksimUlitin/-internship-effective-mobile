package controllers

import (
	"effectiveMobileTask/config"
	"effectiveMobileTask/internal/models"
	"effectiveMobileTask/lib/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

func GetSongDetailAPI(group, song string, c *gin.Context) (models.SongDetail, bool) {
	encodedGroup := url.QueryEscape(group)
	encodedSong := url.QueryEscape(song)

	urlAPI := fmt.Sprintf("%s%s?group=%s&song=%s",
		config.AppConfig.ExternalAPI.BaseURL,
		config.AppConfig.ExternalAPI.InfoURL,
		encodedGroup,
		encodedSong)

	resp, err := http.Get(urlAPI)
	if err != nil {
		return handleAPIError(c, "failed to get song detail", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return handleAPIError(c, "failed to get song detail with status code", resp.StatusCode)
	}

	var dataAPI models.SongDetail
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return handleAPIError(c, "failed to read song detail", err)
	}

	if err := json.Unmarshal(body, &dataAPI); err != nil {
		return handleAPIError(c, "failed to unmarshal song detail", err)
	}

	return dataAPI, false
}

func handleAPIError(c *gin.Context, message string, err interface{}) (models.SongDetail, bool) {
	logger.Error(message, slog.Any("error", err))
	c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	return models.SongDetail{}, true
}

func GetSongDetailJSON(group, song string) (models.SongDetail, error) {
	jsonFileEnrich, err := os.Open("enrichInfoSong.json")
	if err != nil {
		logger.Error("could not open song enrichment file", slog.Any("error", err))
		return models.SongDetail{}, err
	}
	defer jsonFileEnrich.Close()

	jsonVal, err := io.ReadAll(jsonFileEnrich)

	if err != nil {
		logger.Error("failed to read song enrichment file", slog.Any("error", err))
		return models.SongDetail{}, err
	}

	var enrichmentData SongEnriched
	if err := json.Unmarshal(jsonVal, &enrichmentData); err != nil {
		logger.Error("failed to unmarshal song enrichment file", slog.Any("error", err))
		return models.SongDetail{}, err
	}

	if enrichmentData.Group == group && enrichmentData.Song == song {
		return models.SongDetail{
			ReleaseDate: enrichmentData.ReleaseDate,
			Text:        enrichmentData.Text,
			Link:        enrichmentData.Link,
		}, nil
	}

	return models.SongDetail{}, errors.New("invalid song enrichment")
}

func SongEnrichFromJSON(songDetail *models.SongDetail, group, song string) {
	jsonFileEnrich, err := os.Open("enrichInfoSong.json")
	if err != nil {
		logger.Error("failed to open enrichInfoSong.json", slog.Any("error", err))
		return
	}
	defer jsonFileEnrich.Close()

	jsonVal, err := io.ReadAll(jsonFileEnrich)

	if err != nil {
		logger.Error("failed to read enrichInfoSong.json", slog.Any("error", err))
	}

	var enrichmentData SongEnriched
	if err := json.Unmarshal(jsonVal, &enrichmentData); err != nil {
		logger.Error("failed to unmarshal enrichInfoSong.json", slog.Any("error", err))

	} else {
		if enrichmentData.Group == group && enrichmentData.Song == song {
			songDetail.ReleaseDate = enrichmentData.ReleaseDate
			songDetail.Text = enrichmentData.Text
			songDetail.Link = enrichmentData.Link
		}
	}
}
