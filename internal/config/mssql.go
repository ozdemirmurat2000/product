package config

import (
	"fmt"
	"net/url"

	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

func InitDB() gorm.DB {

	encodedPassword := url.QueryEscape(Config.DBPassword)

	host := Config.DBHost
	user := Config.DBUser
	port := Config.DBPort
	dbName := Config.DBName

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, encodedPassword, host, port, dbName)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return *db
}
