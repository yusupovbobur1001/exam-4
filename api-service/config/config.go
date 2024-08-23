package config

import (
	"log"
	"os"

	"github.com/spf13/cast"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP_PORT            string
	AUTH_SERVICE_PORT    string
	BOOKING_SERVICE_PROT string
	SIGNING_KEY          string
	REFRESH_SIGNING_KEY  string
	KafkaBrokers         string
	KafkaTopic           string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env ", err)
	}

	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":2020"))      //
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", "auth-services3:2005"))   //auth-services1
	cfg.BOOKING_SERVICE_PROT = cast.ToString(coalesce("BOOKING_SERVICE_PROT", "booking_service1:6666"))  //booking_service
	cfg.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "secret"))
	cfg.REFRESH_SIGNING_KEY = cast.ToString(coalesce("REFRESH_SIGNING_KEY", "secret1"))
	cfg.KafkaBrokers = cast.ToString(coalesce("KAFKA_BROKERS", "9092"))   //kafka1:9092
	cfg.KafkaTopic = cast.ToString(coalesce("KAFKA_TOPIC", "orders"))
	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
