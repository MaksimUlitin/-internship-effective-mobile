package database

import (
	"effectiveMobileTask/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func DbConnect() *gorm.DB {
	once.Do(func() {
		dsn := config.AppConfig.DB.URL
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to the DB: ", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("failed to get sqlDB from gormDB: ", err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(50)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return db
}
