package config

import (
	"fmt"
	"net/url"
	"productApp/pkg/models"
	"time"

	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

func InitDB() gorm.DB {

	encodedPassword := url.QueryEscape(Config.DBPassword)

	host := Config.DBHost
	user := Config.DBUser
	port := Config.DBPort
	dbName := Config.DBName

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=30&dial+timeout=30",
		user, encodedPassword, host, port, dbName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		DefaultContextTimeout: 60 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.ModelResimModel{})

	return *db
}
