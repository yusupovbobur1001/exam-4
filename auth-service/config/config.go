package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT           string
	USER_SERVICE        string
	USER_ROUTER         string
	DB_HOST             string
	DB_PORT             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
	SIGNING_KEY         string
	REFRESH_SIGNING_KEY string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", "auth-services1"))     //auth-services1
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "postgres-db2"))    //postgres-db2
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "pass"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "on_demand"))
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", "auth-services1:8081"))      // auth-services1:8081
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "secret"))
	config.REFRESH_SIGNING_KEY = cast.ToString(Coalesce("REFRESH_SIGNING_KEY", "secret1"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
