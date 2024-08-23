package mongodb

import (
	"booking_service/config"
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"testing"
)

func TestCreatePayment(t *testing.T) {
	cfg := config.Load()

	mclient, db, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(db)

	req := pb.CreatePaymentRequest{
		BookingId:     "66b892f5d25a97c81486a36b",
		Amount:        2,
		Status:        "rrrrrrrrr",
		PaymentMethod: "rrr",
		TransactionId: "dasdfas",
	}

	r, err := repo.CreatePayment(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestGetPayment(t *testing.T) {
	cfg := config.Load()

	mclient, db, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()


	repo := NewBookingRepo(db)

	req := pb.GetPaymentRequest{
		XId: "66b8ba4f5782ab9b9336253e",
	}

	r, err := repo.GetPayment(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestListPayments(t *testing.T) {
	cfg := config.Load()

	mclient, db, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(db)

	req := pb.ListPaymentsRequest{
		Limit:  10,
		Offset: 0,
	}

	r, err := repo.ListPayments(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}
