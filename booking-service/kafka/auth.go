package kafka

import (
	"booking_service/genproto/booking"
	"booking_service/service"
	"context"
	"encoding/json"
	"log"
)

func BookingRegister(Booking *service.BookingService) func(message []byte) {
	return func(message []byte) {
		var eval booking.CreateBookingRequest
		if err := json.Unmarshal(message, &eval); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}
		respEval, err := Booking.CreateBooking(context.Background(), &eval)
		if err != nil {
			log.Printf("Cannot user register via Kafka: %v", err)
			return
		}
		log.Printf("Register user via Kafka: %+v", respEval)
	}
}

func UpdateBooking(Booking *service.BookingService) func(message []byte) {
	return func(message []byte) {
		var eval booking.UpdateBookingRequest
		if err := json.Unmarshal(message, &eval); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}
		respEval, err := Booking.UpdateBooking(context.Background(), &eval)
		if err != nil {
			log.Printf("Cannot user register via Kafka: %v", err)
			return
		}
		log.Printf("Register user via Kafka: %+v", respEval)
	}
}

func DeleteBooking(Booking *service.BookingService) func(message []byte) {
	return func(message []byte) {
		var eval booking.CancelBookingRequest
		if err := json.Unmarshal(message, &eval); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}
		respEval, err := Booking.CancelBooking(context.Background(), &eval)
		if err != nil {
			log.Printf("Cannot user register via Kafka: %v", err)
			return
		}
		log.Printf("Register user via Kafka: %+v", respEval)
	}
}
