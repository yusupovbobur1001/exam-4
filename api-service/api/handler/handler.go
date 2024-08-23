package handler

import (
	"api_service/config"
	pbb "api_service/genproto/booking"
	pbu "api_service/genproto/user"
	"api_service/pkg"
	"api_service/pkg/kafka"
	"api_service/pkg/logger"
	"log"
	"log/slog"
)

type Handler struct {
	ClientBooking pbb.BookingClient
	ClientUser    pbu.AuthClient
	Logger        *slog.Logger
	Kafka         kafka.ProducerIkafka
}

func NewHandler(cfg *config.Config, kafka kafka.ProducerIkafka) *Handler {
	l, err := logger.NewLoger()
	if err != nil {
		log.Fatal("err: ", err)
	}

	return &Handler{
		ClientBooking: pkg.NewBookingClient(cfg),
		ClientUser:    pkg.NewAuthenticationClient(cfg),
		Logger:        l,
		Kafka:  	   kafka,
	}
}
