package mongodb

import (
	"booking_service/config"
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"testing"
)

func TestCreateReview(t *testing.T) {
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

	req := pb.CreateReviewRequest{
		BookingId:   "66b892f5d25a97c81486a36b",
		UserId:      "734e927e-171e-4efc-b309-c76d4be06e93",
		ProcviderId: "adsfdsgf",
		Rating:      8,
		Comment:     "dsafads",
	}

	r, err := repo.CreateReview(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestUpdateReview(t *testing.T) {
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

	req := pb.UpdateReviewRequest{
		XId:         "66b8be789ee8b6fbf32e1454",
		UserId:      "4afff46f-d335-48f7-984e-b38eafcf60b8",
		ProcviderId: "brbrbrbrbrbr",
		Rating:      5,
		Comment:     "nima gap",
	}

	r, err := repo.UpdateReview(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestDeleteReview(t *testing.T) {
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

	req := pb.DeleteReviewRequest{
		XId: "66b8c0c408a5142fe79a9e1e",
	}

	r, err := repo.DeleteReview(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestListReview(t *testing.T) {
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

	
	req := pb.ListReviewsRequest{
		Limit:  5,
		Offset: 0,
	}

	r, err := repo.ListReviews(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}