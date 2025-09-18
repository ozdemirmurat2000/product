package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigModel struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecretKey string
	RedisHost    string
	RedisPort    string
}

var Config ConfigModel

func InitConfig() {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	Config.DBHost = os.Getenv("DB_HOST")
	Config.DBPort = os.Getenv("DB_PORT")
	Config.DBPassword = os.Getenv("DB_PASSWORD")
	Config.DBName = os.Getenv("DB_NAME")
	Config.DBUser = os.Getenv("DB_USER")
	Config.JWTSecretKey = os.Getenv("JWT_SECRET")
	Config.RedisHost = os.Getenv("REDIS_HOST")
	Config.RedisPort = os.Getenv("REDIS_PORT")

}
