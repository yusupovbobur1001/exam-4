package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BookingService string
	MongoURI       string
	MongoDBName    string
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.BookingService = cast.ToString(Coalesce("BOOKING_SERVICE", ":6666"))
	config.MongoURI = cast.ToString(Coalesce("MONGO_URI", "mongodb://mongo:27018"))     //mongodb://mongosh:27017
	config.MongoDBName = cast.ToString(Coalesce("MONGODB_NAME", "booking_service1"))
	config.RedisAddr = cast.ToString(Coalesce("REDIS_ADDR", ":6379"))    //redis:6379
	config.RedisPassword = cast.ToString(Coalesce("REDIS_PASSWORD", ""))
	config.RedisDB = cast.ToInt(Coalesce("REDIS_DB", 0))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
