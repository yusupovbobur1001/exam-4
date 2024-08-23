package mongodb

import (
	"booking_service/config"
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	// cfg := config.Load()

	// mclient, mongodb, err := NewMongoClient(&cfg)
	// if err != nil {
	// 	fmt.Println(1)
	// 	panic(err)
	// }

	// defer func() {
	// 	if err := mclient.Disconnect(context.Background()); err != nil {
	// 		t.Fatal(err)
	// 	}
	// }()

	// repo := NewBookingRepo(mongodb)
	
	// req := pb.CreateBookingRequest{
	// 	UserId:        "adskjlfns",
	// 	ProviderId:    "sdakl",
	// 	ServiceId:     "asdk",
	// 	Status:        "sakdlm",
	// 	ScheduledTime: &pb.ScheduledTime{
	// 		StartTime: "asdf",
	// 		EndTime:   "asdfa",
	// 	},
	// 	TatolPrice:    0,
	// 	Location:      &pb.Location{
	// 		City:    "dsf",
	// 		Country: "sadf",
	// 	},
	// }

	// r, err := repo.CreateBooking(context.Background(), &req)

	// assert.NoError(t, err)
	// assert.Equal(t, req.ProviderId, r.ProviderId)
	// assert.Equal(t, req.ServiceId, r.ServiceId)
	// assert.Equal(t, req.Location.City, r.Location.City)
	// assert.Equal(t, req.Location.Country, r.Location.Country)
	// assert.Equal(t, req.Status, r.Status)
	// assert.Equal(t, req.UserId, r.UserId)
}

func TestGetBooking(t *testing.T) {
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

	req := pb.GetBookingRequest{
		XId: "66b88d65e06b6b31ff43cc80",
	}

	result, err := repo.GetBooking(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, req.XId, result.XId)
	fmt.Println(result)
}

func TestUpdateBooking(t *testing.T) {
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

	req := pb.UpdateBookingRequest{
		UserId:     "ddddd",
		ProviderId: "dfdfg",
		ServiceId:  "ffff",
		Status:     "fffff",
		TatolPrice: 1,
		XId:        "66b88d65e06b6b31ff43cc80",
	}

	r, err := repo.UpdateBooking(context.Background(), &req)
	if err !=  nil {
		t.Fatal(err)
	}

	assert.Equal(t, req.UserId, r.UserId)
	assert.Equal(t, req.ProviderId, r.ProviderId)
	fmt.Println(r)
}

func TestCancelBooking(t *testing.T) {
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

	req := pb.CancelBookingRequest{
		XId: "66b89299c1560cbe10317905",
	}

	r, err := repo.CancelBooking(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "SUCCESS", r.Message)
}

func TestListBookings(t *testing.T) {
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

	req := pb.ListBookingsRequest{
		Limit:  10,
		Offset: 0,
	}

	r, err := repo.ListBookings(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(r)
}

func TestGetMostFrequentServiceID1(t *testing.T) {
	cfg := config.Load()

	mclient, db, err := NewMongoClient(&cfg)
	if err != nil {
		t.Log(err)
	}

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	repo := NewBookingRepo(db)

	s, err := repo.GetMostFrequentServiceID(context.Background(), &pb.Void{})
	if err != nil {
		t.Log(err)
	}

	fmt.Println(s, )
}



