package mongodb

import (
	"booking_service/config"
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"testing"
)

func TestCreateCreateProviders(t *testing.T) {
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

	req := pb.CreateProvidersRequest{
		UserId:      "734e927e-171e-4efc-b309-c76d4be06e93",
		CompanyName: "nimadir",
		ServiceId: []*pb.ServiceId{
			{
				XId: "734e927e-171e-4efc-b309-c76d4be06e93", 
			},
		},
		Location: &pb.Location{
			City:    "Tashkent", 
			Country: "Uzbekistan", 
		},
		Availabilitys: []*pb.AvailabilityR{
			{
				StartTime: "09:00", 
				EndTime:   "17:00", 
			},
			{
				StartTime: "18:00",
				EndTime:   "21:00",
			},
		},
	}

	r, err := repo.CreateProviders(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	t.Log(r.XId)
}

func TestUpdateProviders(t *testing.T) {
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

	req := pb.UpdateProvidersRequest{
		UserId:        "734e927e-171e-4efc-b309-c76d4be06e93", // Example user ID
		CompanyName:   "Updated Company Name", // New company name
		Location: &pb.Location{
			City:    "Samarkand", // Updated city
			Country: "Uzbekistan", // Updated country
		},
		Availabilitys: []*pb.AvailabilityR{
			{
				StartTime: "08:00", // Updated start time
				EndTime:   "17:00", // Updated end time
			},
			{
				StartTime: "19:00", // Another availability slot
				EndTime:   "22:00",
			},
		},
		XId:           "66b8ca371f33a7ce5a4799dd", // ID of the provider to update
	}
	
	r, err := repo.UpdateProviders(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestDeleteProviders(t *testing.T) {
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

	req := pb.DeleteProvidersRequest{
		XId: "66b8ca371f33a7ce5a4799dd",
	}

	r, err := repo.DeleteProviders(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestGetProviders(t *testing.T) {
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

	req := pb.GetProvidersRequest{
		XId: "66b8cb55a6ebbc466fe73761",
	}

	r, err := repo.GetProviders(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestListProviders(t *testing.T) {
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

	req := pb.ListProvidersRequest{
		Limit:  2,
		Offset: 0,
	}

	r, err := repo.ListProviders(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(r)
	fmt.Println(r)
}

func TestSearchProviders(t *testing.T) {
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

	req := pb.SearchProvidersRequest{
		UserId:        "734e927e-171e-4efc-b309-c76d4be06e93",
		CompanyName:   "nimadir",
		Availabilitys: []*pb.AvailabilityR{},
		Location:      &pb.Location{},
	}

	r, err := repo.SearchProviders(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Provider)
}
