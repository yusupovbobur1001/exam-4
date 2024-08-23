package pkg

import (
	"api_service/config"
	pbuBooking "api_service/genproto/booking"
	pbuAuth "api_service/genproto/user"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthenticationClient(cfg *config.Config) pbuAuth.AuthClient {
	conn, err := grpc.Dial(cfg.AUTH_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println(cfg.AUTH_SERVICE_PORT)
	if err != nil {
		fmt.Println(err)
		log.Println("error while connecting authentication service ", err)
	}

	return pbuAuth.NewAuthClient(conn)

}

func NewBookingClient(cfg *config.Config) pbuBooking.BookingClient {
	fmt.Println(cfg.BOOKING_SERVICE_PROT)
	conn, err := grpc.Dial(cfg.BOOKING_SERVICE_PROT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}

	return pbuBooking.NewBookingClient(conn)
}
