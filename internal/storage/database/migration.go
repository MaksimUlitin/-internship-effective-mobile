package database

import (
	"effectiveMobileTask/internal/models"
	"effectiveMobileTask/lib/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Group{}, &models.Song{}); err != nil {
		logger.Error("Database migration failed", "error", err)
		return err
	}
	logger.Info("Database migration completed successfully")

	return nil
}
