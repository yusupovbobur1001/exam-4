package mongodb

import (
	pb "booking_service/genproto/booking"
	"booking_service/storage/redis"
	"booking_service/storage/repo"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookingRepo struct {
	db *mongo.Database
}

func NewBookingRepo(db *mongo.Database) repo.StorageInterfase {
	return &BookingRepo{db: db}
}

func (b *BookingRepo) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	collection := b.db.Collection("booking")
	fmt.Println(req, "00000000000000000000000000000000")
	booking := bson.M{
		"user_id":     req.GetUserId(),
		"provider_id": req.GetProviderId(),
		"service_id":  req.GetServiceId(),
		"status":      req.GetStatus(),
		"scheduled_time": bson.M{
			"start_time": req.GetScheduledTime().GetStartTime(),
			"end_time":   req.GetScheduledTime().GetEndTime(),
		},
		"tatol_price": req.GetTatolPrice(),
		"location": bson.M{
			"city":    req.GetLocation().GetCity(),
			"country": req.GetLocation().GetCountry(),
		},
	}

	resp, err := collection.InsertOne(ctx, booking)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}

	oid, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		fmt.Println(resp.InsertedID)
		return nil, fmt.Errorf("failed to convert objectid to hex: %v", resp.InsertedID)
	}
	fmt.Println(resp.InsertedID)
	return &pb.CreateBookingResponse{
		XId:        oid.Hex(),
		UserId:     req.UserId,
		ProviderId: req.ProviderId,
		ServiceId:  req.ServiceId,
		Status:     req.GetStatus(),
		TatolPrice: req.TatolPrice,
		Location: &pb.Location{
			City:    req.GetLocation().GetCity(),
			Country: req.GetLocation().GetCountry(),
		},
	}, nil
}

func (b *BookingRepo) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	collection := b.db.Collection("booking")

	oid, err := primitive.ObjectIDFromHex(req.XId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}

	var booking struct {
		UserId        string `bson:"user_id"`
		ProviderId    string `bson:"provider_id"`
		ServiceId     string `bson:"service_id"`
		Status        string `bson:"status"`
		ScheduledTime struct {
			StartTime string `bson:"start_time"`
			EndTime   string `bson:"end_time"`
		} `bson:"scheduled_time"`
		TatolPrice float64 `bson:"tatol_price"`
		Location   struct {
			City    string `bson:"city"`
			Country string `bson:"country"`
		} `bson:"location"`
		Id primitive.ObjectID `bson:"_id"`
	}

	err = collection.FindOne(ctx, filter).Decode(&booking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("booking not found")
		}
		return nil, err
	}

	resp := &pb.GetBookingResponse{
		UserId:     booking.UserId,
		ProviderId: booking.ProviderId,
		ServiceId:  booking.ServiceId,
		Status:     booking.Status,
		ScheduledTime: &pb.ScheduledTime{
			StartTime: booking.ScheduledTime.StartTime,
			EndTime:   booking.ScheduledTime.EndTime,
		},
		TatolPrice: float32(booking.TatolPrice),
		Location: &pb.Location{
			City:    booking.Location.City,
			Country: booking.Location.Country,
		},
		XId: booking.Id.Hex(),
	}

	return resp, nil
}

func (b *BookingRepo) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	fmt.Println("11111111111111111111111111111111111111111111111111111111111")
	collection := b.db.Collection("booking")

	oid, err := primitive.ObjectIDFromHex(req.XId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": oid,
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":     req.GetUserId(),
			"provider_id": req.GetProviderId(),
			"service_id":  req.GetServiceId(),
			"status":      req.GetStatus(),
			"tatol_price": req.GetTatolPrice(),
			"updated_at":  time.Now().Format("2006-01-02 15:04:05"), // Yang
		},
	}
	fmt.Println(req.ProviderId, req.ServiceId, req.UserId, req.XId, "\n", "----+-+-+-+-++--+-+-++---++--")
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update booking: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("booking not found")
	}

	var updatedBooking struct {
		UserId     string             `bson:"user_id"`
		ProviderId string             `bson:"provider_id"`
		ServiceId  string             `bson:"service_id"`
		Status     string             `bson:"status"`
		TatolPrice float64            `bson:"tatol_price"`
		Id         primitive.ObjectID `bson:"_id"`
		Updated_at string             `bson:"updated_at"`
	}

	err = collection.FindOne(ctx, filter).Decode(&updatedBooking)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated booking: %v", err)
	}

	return &pb.UpdateBookingResponse{
		UserId:     updatedBooking.UserId,
		ProviderId: updatedBooking.ProviderId,
		ServiceId:  updatedBooking.ServiceId,
		Status:     updatedBooking.Status,
		TatolPrice: float32(updatedBooking.TatolPrice),
		UpdatedAt:  updatedBooking.Updated_at,
	}, nil
}

func (b *BookingRepo) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error) {
	collection := b.db.Collection("booking")

	oid, err := primitive.ObjectIDFromHex(req.XId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, err
	}

	return &pb.CancelBookingResponse{
		Message: "SUCCESS",
	}, nil
}

func (b *BookingRepo) ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	collection := b.db.Collection("booking")
	fmt.Println(req, "1111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	// MongoDB qidiruv parametrlari
	findOptions := options.Find()
	findOptions.SetLimit(int64(req.GetLimit()))
	findOptions.SetSkip(int64(req.GetOffset()))

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to list bookings: %v", err)
	}
	defer cursor.Close(ctx)

	var bookings []*pb.GetBookingR
	fmt.Println(req, "2222222222222222222222222222222222222222222")
	for cursor.Next(ctx) {
		var booking struct {
			UserId        string `bson:"user_id"`
			ProviderId    string `bson:"provider_id"`
			ServiceId     string `bson:"service_id"`
			Status        string `bson:"status"`
			ScheduledTime struct {
				StartTime string `bson:"start_time"`
				EndTime   string `bson:"end_time"`
			} `bson:"scheduled_time"`
			TatolPrice float64 `bson:"tatol_price"`
			Location   struct {
				City    string `bson:"city"`
				Country string `bson:"country"`
			} `bson:"location"`
			Id primitive.ObjectID `bson:"_id"`
		}

		err := cursor.Decode(&booking)
		if err != nil {
			return nil, fmt.Errorf("failed to decode booking: %v", err)
		}

		// Protobuf GetBookingR message'iga o'tkazish
		bookings = append(bookings, &pb.GetBookingR{
			UserId:     booking.UserId,
			ProviderId: booking.ProviderId,
			ServiceId:  booking.ServiceId,
			Status:     booking.Status,
			ScheduledTime: &pb.ScheduledTime{
				StartTime: booking.ScheduledTime.StartTime,
				EndTime:   booking.ScheduledTime.EndTime,
			},
			TatolPrice: float32(booking.TatolPrice),
			Location: &pb.Location{
				City:    booking.Location.City,
				Country: booking.Location.Country,
			},
			XId: booking.Id.Hex(),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	fmt.Println(bookings, "33333333333333333333333333333")
	return &pb.ListBookingsResponse{
		Listbooks: bookings,
	}, nil
}

func (b *BookingRepo) GetMostFrequentServiceID(ctx context.Context, req *pb.Void) (*pb.GetMostRequest, error) {
	coll := b.db.Collection("booking")

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":   "$service_id",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{
				"count": -1,
			},
		},
		{
			"$limit": 1,
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("Aggregationda xato:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var result struct {
		ServiceID string `bson:"_id"`
		Count     int32  `bson:"count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("natija topilmadi")
	}

	rdb := redis.RedisConn()
	err = rdb.HSet(ctx, "papular-service", "service_id", result.ServiceID).Err()
	if err != nil {
		return nil, err
	}
	err = rdb.HSet(ctx, "papular-service", "count", result.Count).Err()
	if err != nil {
		return nil, err
	}

	p := pb.List{
		ServiceId: result.ServiceID,
		Count:     result.Count,
	}

	resp := pb.GetMostRequest{}
	resp.Body = append(resp.Body, &p)

	return &resp, nil
}

