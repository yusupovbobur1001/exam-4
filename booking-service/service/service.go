package service

import (
	pb "booking_service/genproto/booking"
	"booking_service/storage/redis"
	"booking_service/storage/repo"
	"context"
	"fmt"
	"log/slog"
	"strconv"
)

type BookingManagementService interface {
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
}

type BookingService struct {
	pb.UnimplementedBookingServer
	Logger      *slog.Logger
	bookingRepo repo.StorageInterfase
}

func NewBookingService(logger *slog.Logger, bookingRepo repo.StorageInterfase) *BookingService {
	return &BookingService{
		Logger:      logger,
		bookingRepo: bookingRepo,
	}
}

func (b *BookingService) GetMostFrequentServiceID(ctx context.Context, req *pb.Void) (*pb.GetMostRequest, error) {
	rdb := redis.RedisConn()
	fmt.Println("00000000000000000000000")
	result, err := rdb.HGetAll(ctx, "papular-service").Result()
	if len(result) == 0 || err != nil {
		fmt.Println("2345678765434567")
		res, err := b.bookingRepo.GetMostFrequentServiceID(ctx, &pb.Void{})
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	resps := pb.GetMostRequest{}
	serviceID, ok1 := result["service_id"]
	countStr, ok2 := result["count"]
	
	if ok1 && ok2 {
		r := &pb.List{}
		c, _ := strconv.Atoi(countStr)
		r.Count = int32(c)
		r.ServiceId = serviceID
		resps.Body = append(resps.Body, r)
	}

	fmt.Println("tttttttttttttttttttt")
	return &resps, nil

}

func (b *BookingService) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	resp, err := b.bookingRepo.CreateBooking(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to create booking", "error", err)
		return nil, err
	}
	fmt.Println(resp, "--------------")
	return resp, nil
}

func (b *BookingService) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	resp, err := b.bookingRepo.GetBooking(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to get booking", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	fmt.Println(req, "\n", "---------------+-+------------------")
	resp, err := b.bookingRepo.UpdateBooking(ctx, req)
	if err != nil {
		fmt.Println(err)
		b.Logger.Error("Failed to update booking", "error", err)
		return nil, err
	}
	fmt.Println(resp)
	return resp, nil

}

func (b *BookingService) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error) {
	resp, err := b.bookingRepo.CancelBooking(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to cancel booking", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	resp, err := b.bookingRepo.ListBookings(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to list bookings", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) SearchServices(ctx context.Context, req *pb.SearchServicesRequest) (*pb.SearchServicesResponse, error) {
	fmt.Println(req)
	resp, err := b.bookingRepo.SearchServices(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to search services", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	resp, err := b.bookingRepo.CreateService(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to create service", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*pb.UpdateServiceResponse, error) {
	resp, err := b.bookingRepo.UpdateService(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to update service", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*pb.DeleteServiceResponse, error) {
	resp, err := b.bookingRepo.DeleteService(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to delete service", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	resp, err := b.bookingRepo.ListServices(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to list services", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	resp, err := b.bookingRepo.CreatePayment(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to create payment", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	resp, err := b.bookingRepo.GetPayment(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to get payment", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error) {
	resp, err := b.bookingRepo.ListPayments(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to list payments", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	resp, err := b.bookingRepo.CreateReview(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to create review", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error) {
	fmt.Println(req)
	resp, err := b.bookingRepo.UpdateReview(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to update review", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	resp, err := b.bookingRepo.DeleteReview(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to delete review", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	resp, err := b.bookingRepo.ListReviews(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to list reviews", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) CreateProviders(ctx context.Context, req *pb.CreateProvidersRequest) (*pb.CreateProvidersResponse, error) {
	resp, err := b.bookingRepo.CreateProviders(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to create providers", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) UpdateProviders(ctx context.Context, req *pb.UpdateProvidersRequest) (*pb.UpdateProvidersResponse, error) {
	resp, err := b.bookingRepo.UpdateProviders(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to create providers", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) DeleteProviders(ctx context.Context, req *pb.DeleteProvidersRequest) (*pb.DeleteProvidersResponse, error) {
	resp, err := b.bookingRepo.DeleteProviders(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to create providers", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) GetProviders(ctx context.Context, req *pb.GetProvidersRequest) (*pb.GetProvidersResponse, error) {
	fmt.Println(req.XId)
	resp, err := b.bookingRepo.GetProviders(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to get providers", "error", err)
		return nil, err
	}
	fmt.Println(resp, "-----------")
	return resp, nil
}

func (b *BookingService) ListProviders(ctx context.Context, req *pb.ListProvidersRequest) (*pb.ListProvidersResponse, error) {
	resp, err := b.bookingRepo.ListProviders(ctx, req)
	if err != nil {
		b.Logger.Error("Failed to list providers", "error", err)
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) SearchProviders(ctx context.Context, req *pb.SearchProvidersRequest) (*pb.SearchProvidersResponse, error) {
	res, err := b.bookingRepo.SearchProviders(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
