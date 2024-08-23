package main

import (
	"booking_service/config"
	"booking_service/kafka"
	"booking_service/logger"
	"booking_service/service"
	"booking_service/storage/mongodb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "booking_service/genproto/booking"
)

func main() {
	cfg := config.Load()

	listener, err := net.Listen("tcp", cfg.BookingService)
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	logger := logger.NewLogger()

	mongoClient, db, err := mongodb.NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	bookingRepo := mongodb.NewBookingRepo(db)
	bookingService := service.NewBookingService(logger, bookingRepo)

	brokers := []string{"kafka1:9093"}   //kafka1:9092 
	fmt.Println(brokers)
	kcm := kafka.NewKafkaConsumerManager()
	if err := kcm.RegisterConsumer(brokers, "createBooking", "register", kafka.BookingRegister(bookingService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'createBooking' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	if err := kcm.RegisterConsumer(brokers, "updateBooking", "update", kafka.UpdateBooking(bookingService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'updateBooking' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "deleteBooking", "delete", kafka.DeleteBooking(bookingService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'deleteBooking' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookingServer(grpcServer, bookingService)

	log.Printf("Server is listening on port %s\n", cfg.BookingService)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
