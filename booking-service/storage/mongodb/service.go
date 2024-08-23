package mongodb

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BookingRepo) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	collection := s.db.Collection("services")

	// Yangi xizmatni yaratish
	service := bson.M{
		"user_id":       req.GetUserId(),
		"descrioptions": req.GetDescrioptions(),
		"duration":      req.GetDuration(),
		"price":         req.GetPrice(),
	}

	// MongoDB'ga qo'shish
	resp, err := collection.InsertOne(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("failed to insert service: %v", err)
	}

	// ID'ni olish
	oid, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID: %v", resp.InsertedID)
	}

	// Javobni tayyorlash
	return &pb.CreateServiceResponse{
		UserId:        req.GetUserId(),
		Descrioptions: req.GetDescrioptions(),
		Duration:      req.GetDuration(),
		Price:         req.GetPrice(),
		XId:           oid.Hex(),
	}, nil
}

func (s *BookingRepo) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*pb.UpdateServiceResponse, error) {
	collection := s.db.Collection("services")

	// ID'ni olish va formatlash
	oid, err := primitive.ObjectIDFromHex(req.XId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to ObjectID: %v", err)
	}

	// Filtr va yangilanish ma'lumotlarini tayyorlash
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"user_id":       req.GetUserId(),
			"price":         req.GetPrice(),
			"duration":      req.GetDuration(),
			"descrioptions": req.GetDescrioptions(),
			"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	// Yangilash
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update service: %v", err)
	}

	// Agar hech narsa yangilanmagan bo'lsa
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("service not found")
	}

	// Yangilangan xizmatni olish
	var updatedService struct {
		UserId        string             `bson:"user_id"`
		Price         float64            `bson:"price"`
		Duration      int32              `bson:"duration"`
		Descrioptions string             `bson:"descrioptions"`
		Id            primitive.ObjectID `bson:"_id"`
		UpdatedAt     string             `bson:"updated_at"`
	}

	err = collection.FindOne(ctx, filter).Decode(&updatedService)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated service: %v", err)
	}

	// Javob tayyorlash
	return &pb.UpdateServiceResponse{
		XId:           updatedService.Id.Hex(),
		UserId:        updatedService.UserId,
		Price:         float32(updatedService.Price),
		Duration:      updatedService.Duration,
		Descrioptions: updatedService.Descrioptions,
		UpdatedAt:     updatedService.UpdatedAt,
	}, nil
}

func (s *BookingRepo) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*pb.DeleteServiceResponse, error) {
	collection := s.db.Collection("services")

	// ID'ni olish va formatlash
	oid, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to ObjectID: %v", err)
	}

	// Filtrni tayyorlash
	filter := bson.M{"_id": oid}

	// Xizmatni o'chirish
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete service: %v", err)
	}

	// Agar hech qanday hujjat o'chirilmagan bo'lsa
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("service not found")
	}

	// Javobni qaytarish
	return &pb.DeleteServiceResponse{
		Message: "Service deleted successfully",
	}, nil
}

func (s *BookingRepo) ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	collection := s.db.Collection("services")

	// Limit va offset ni olish
	limit := int64(req.GetLimit())
	offset := int64(req.GetOffset())

	// Ma'lumotlarni olish uchun options
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	// Xizmatlarni olish
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to list services: %v", err)
	}
	defer cursor.Close(ctx)

	var services []*pb.ListServiceR
	for cursor.Next(ctx) {
		var service struct {
			UserId        string             `bson:"user_id"`
			Descrioptions string             `bson:"descrioptions"`
			Duration      int32              `bson:"duration"`
			Price         float64            `bson:"price"`
			Id            primitive.ObjectID `bson:"_id"`
		}

		if err := cursor.Decode(&service); err != nil {
			return nil, fmt.Errorf("failed to decode service: %v", err)
		}

		services = append(services, &pb.ListServiceR{
			UserId:        service.UserId,
			Descrioptions: service.Descrioptions,
			Duration:      service.Duration,
			Price:         float32(service.Price),
			XId:           service.Id.Hex(),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &pb.ListServicesResponse{
		ListServices: services,
	}, nil
}

func (p *BookingRepo) SearchServices(ctx context.Context, req *pb.SearchServicesRequest) (*pb.SearchServicesResponse, error) {
	collection := p.db.Collection("services")

	filter := bson.M{}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}

	if req.Price > 0 {
		filter["price"] = req.Price
	}

	if req.Duration > 0 {
		filter["duration"] = req.Duration
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Servislarni qidirishda xatolik: %v", err)
	}
	fmt.Println(cursor, filter)
	defer cursor.Close(ctx)

	var services []*pb.SearchServiceR
	for cursor.Next(ctx) {
		var service struct {
			UserId        string  `bson:"user_id"`
			Descrioptions string  `bson:"descrioptions"`
			Duration      int32   `bson:"duration"`
			Price         float64 `bson:"price"`
			XId           string  `bson:"_id"`
		}
		if err := cursor.Decode(&service); err != nil {
			return nil, status.Errorf(codes.Internal, "Kursorni dekodlashda xatolik: %v", err)
		}
		services = append(services, &pb.SearchServiceR{
			UserId:        service.UserId,
			Descrioptions: service.Descrioptions,
			Duration:      service.Duration,
			Price:         float32(service.Price),
			XId:           service.XId,
		})
	}
	if err := cursor.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "Kursorni olishda xatolik: %v", err)
	}
	if len(services) == 0 {
		return nil, status.Errorf(codes.NotFound, "Hech qanday mos keluvchi hujjat topilmadi")
	}
	fmt.Println("Natijalar:", services)

	return &pb.SearchServicesResponse{
		Services: services,
	}, nil
}
