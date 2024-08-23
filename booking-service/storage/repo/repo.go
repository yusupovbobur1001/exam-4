package repo

import (
	"context"

	pb "booking_service/genproto/booking"
)

type StorageInterfase interface {
	CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error)
	GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error)
	UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error)
	CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error)
	ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error)
	CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error)
	UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*pb.UpdateServiceResponse, error)
	DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*pb.DeleteServiceResponse, error)
	ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error)
	CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error)
	GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error)
	ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error)
	CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error)
	UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error)
	DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error)
	ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error)
	CreateProviders(ctx context.Context, req *pb.CreateProvidersRequest) (*pb.CreateProvidersResponse, error)
	UpdateProviders(ctx context.Context, req *pb.UpdateProvidersRequest) (*pb.UpdateProvidersResponse, error)
	DeleteProviders(ctx context.Context, req *pb.DeleteProvidersRequest) (*pb.DeleteProvidersResponse, error)
	GetProviders(ctx context.Context, req *pb.GetProvidersRequest) (*pb.GetProvidersResponse, error)
	ListProviders(ctx context.Context, req *pb.ListProvidersRequest) (*pb.ListProvidersResponse, error)
	SearchProviders(ctx context.Context, req *pb.SearchProvidersRequest) (*pb.SearchProvidersResponse, error)
	SearchServices(ctx context.Context, req *pb.SearchServicesRequest) (*pb.SearchServicesResponse, error)
	GetMostFrequentServiceID(ctx context.Context, req *pb.Void) (*pb.GetMostRequest, error)
}
