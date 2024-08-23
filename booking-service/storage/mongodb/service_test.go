package mongodb

import (
	"booking_service/config"
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateService(t *testing.T) {
	cfg := config.Load()

	mclient, mongodb, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(mongodb)

	req := pb.CreateServiceRequest{
		UserId:        "sjdfas",
		Descrioptions: "asdfas",
		Duration:      2,
		Price:         2,
	}

	r, err := repo.CreateService(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, req.UserId, r.UserId)
	fmt.Println(r)
}

func TestUpdateService(t *testing.T) {
	cfg := config.Load()

	mclient, mongodb, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(mongodb)

	req := pb.UpdateServiceRequest{
		XId:           "66b93bc88b5b33890b634ecb",
		UserId:        "aaaaaaa",
		Price:         1,
		Duration:      2,
		Descrioptions: "bbbbbb",
	}

	r, err := repo.UpdateService(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, req.UserId, r.UserId)
	fmt.Println(r)
}

func TestDeleteService(t *testing.T) {
	cfg := config.Load()

	mclient, mongodb, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(mongodb)

	req := pb.DeleteServiceRequest{
		XId: "66b897b52ceb02c46967e9db",
	}

	r, err := repo.DeleteService(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Service deleted successfully", r.Message)
	fmt.Println(r)
}

func TestListServices(t *testing.T) {
	cfg := config.Load()

	mclient, mongodb, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(mongodb)

	req := pb.ListServicesRequest{
		Limit:  10,
		Offset: 0,
	}

	r, err := repo.ListServices(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.ListServices)
}

func TestSearchServices(t *testing.T) {
	cfg := config.Load()

	mclient, db, err := NewMongoClient(&cfg)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(db)

	req := pb.SearchServicesRequest{
		UserId:   "",
		Price:    2,
		Duration: 0,
	}

	r, err := repo.SearchServices(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
