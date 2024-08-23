package mongodb

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *BookingRepo) CreateProviders(ctx context.Context, req *pb.CreateProvidersRequest) (*pb.CreateProvidersResponse, error) {
	collection := p.db.Collection("providers")

	var serviceIds []primitive.ObjectID

	provider := bson.M{
		"user_id":       req.GetUserId(),
		"company_name":  req.GetCompanyName(),
		"service_id":    serviceIds,
		"location":      req.GetLocation(),
		"availabilitys": req.GetAvailabilitys(),
		"created_at":    time.Now().Format(time.RFC3339),
	}

	result, err := collection.InsertOne(ctx, provider)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateProvidersResponse{
		UserId:        req.GetUserId(),
		CompanyName:   req.GetCompanyName(),
		ServiceId:     req.ServiceId,
		Location:      req.GetLocation(),
		Availabilitys: req.GetAvailabilitys(),
		XId:           result.InsertedID.(primitive.ObjectID).Hex(), // Get the inserted ID as a string
		CreratedAt:    provider["created_at"].(string),
	}

	return resp, nil
}

func (p *BookingRepo) UpdateProviders(ctx context.Context, req *pb.UpdateProvidersRequest) (*pb.UpdateProvidersResponse, error) {
	collection := p.db.Collection("providers")
	
	oid, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":      req.GetUserId(),
			"company_name": req.GetCompanyName(),
			"location": bson.M{
				"city":    req.GetLocation().GetCity(),
				"country": req.GetLocation().GetCountry(),
			},
			"availabilitys": req.GetAvailabilitys(),
			"updated_at":    time.Now().Format(time.RFC3339),
		},
	}

	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updatedProvider struct {
		UserId      string               `bson:"user_id"`
		CompanyName string               `bson:"company_name"`
		ServiceIds  []primitive.ObjectID `bson:"service_id"`
		Location    struct {
			City    string `bson:"city"`
			Country string `bson:"country"`
		} `bson:"location"`
		Availabilitys []struct {
			StartTime string `bson:"start_time"`
			EndTime   string `bson:"end_time"`
		} `bson:"availabilitys"`
		XId       primitive.ObjectID `bson:"_id"`
		CreatedAt string             `bson:"created_at"`
		UpdatedAt string             `bson:"updated_at"`
	}

	err = result.Decode(&updatedProvider)
	if err != nil {
		return nil, err
	}

	var serviceIds []*pb.ServiceId
	for _, sid := range updatedProvider.ServiceIds {
		serviceIds = append(serviceIds, &pb.ServiceId{
			XId: sid.Hex(),
		})
	}

	var availabilitys []*pb.AvailabilityR
	for _, a := range updatedProvider.Availabilitys {
		availabilitys = append(availabilitys, &pb.AvailabilityR{
			StartTime: a.StartTime,
			EndTime:   a.EndTime,
		})
	}

	resp := &pb.UpdateProvidersResponse{
		UserId:      updatedProvider.UserId,
		CompanyName: updatedProvider.CompanyName,
		ServiceId:   serviceIds,
		Location: &pb.Location{
			City:    updatedProvider.Location.City,
			Country: updatedProvider.Location.Country,
		},
		Availabilitys: availabilitys,
		XId:           updatedProvider.XId.Hex(),
		CreratedAt:    updatedProvider.CreatedAt,
		UpdatedAt:     updatedProvider.UpdatedAt,
	}

	return resp, nil
}

func (p *BookingRepo) DeleteProviders(ctx context.Context, req *pb.DeleteProvidersRequest) (*pb.DeleteProvidersResponse, error) {
	collection := p.db.Collection("providers")

	oid, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	filter := bson.M{"_id": oid}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("no provider found with the given ID")
	}

	resp := &pb.DeleteProvidersResponse{
		Message: "Provider successfully deleted",
	}

	return resp, nil
}

func (p *BookingRepo) GetProviders(ctx context.Context, req *pb.GetProvidersRequest) (*pb.GetProvidersResponse, error) {
	collection := p.db.Collection("providers")

	// Convert the string ID to ObjectID
	oid, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// Create the filter to find the provider
	filter := bson.M{"_id": oid}

	// Define a struct to hold the provider data
	var provider struct {
		UserId      string               `bson:"user_id"`
		CompanyName string               `bson:"company_name"`
		ServiceIds  []primitive.ObjectID `bson:"service_id"`
		Location    struct {
			City    string `bson:"city"`
			Country string `bson:"country"`
		} `bson:"location"`
		Availabilitys []struct {
			StartTime string `bson:"starttime"`
			EndTime   string `bson:"endtime"`
		} `bson:"availabilitys"`
		XId       primitive.ObjectID `bson:"_id"`
		CreatedAt string             `bson:"created_at"`
		UpdatedAt string             `bson:"updated_at"`
	}

	// Find the provider in the database
	err = collection.FindOne(ctx, filter).Decode(&provider)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no provider found with the given ID")
		}
		return nil, err
	}

	var serviceIds []*pb.ServiceId
	for _, sid := range provider.ServiceIds {
		serviceIds = append(serviceIds, &pb.ServiceId{
			XId: sid.Hex(),
		})
	}

	resp := &pb.GetProvidersResponse{
		UserId:      provider.UserId,
		CompanyName: provider.CompanyName,
		ServiceId:   serviceIds,
		Location: &pb.Location{
			City:    provider.Location.City,
			Country: provider.Location.Country,
		},
		Availabilitys: func() []*pb.AvailabilityR {
			availabilitys := make([]*pb.AvailabilityR, len(provider.Availabilitys))
			for i, a := range provider.Availabilitys {
				availabilitys[i] = &pb.AvailabilityR{
					StartTime: a.StartTime,
					EndTime:   a.EndTime,
				}
			}
			return availabilitys
		}(),
		XId:        provider.XId.Hex(),
		CreratedAt: provider.CreatedAt,
		UpdatedAt:  provider.UpdatedAt,
	}

	return resp, nil
}

func (p *BookingRepo) ListProviders(ctx context.Context, req *pb.ListProvidersRequest) (*pb.ListProvidersResponse, error) {
	collection := p.db.Collection("providers")

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var providers []struct {
		UserId      string               `bson:"user_id"`
		CompanyName string               `bson:"company_name"`
		ServiceIds  []primitive.ObjectID `bson:"service_id"`
		Location    struct {
			City    string `bson:"city"`
			Country string `bson:"country"`
		} `bson:"location"`
		Availabilitys []struct {
			StartTime string `bson:"starttime"`
			EndTime   string `bson:"endtime"`
		} `bson:"availabilitys"`
		XId       primitive.ObjectID `bson:"_id"`
		CreatedAt string             `bson:"created_at"`
	}

	if err := cursor.All(ctx, &providers); err != nil {
		return nil, err
	}

	var responseProviders []*pb.ListProviderR
	for _, provider := range providers {
		var serviceIds []*pb.ServiceId
		for _, sid := range provider.ServiceIds {
			serviceIds = append(serviceIds, &pb.ServiceId{
				XId: sid.Hex(),
			})
		}

		var availabilitys []*pb.AvailabilityR
		for _, a := range provider.Availabilitys {
			availabilitys = append(availabilitys, &pb.AvailabilityR{
				StartTime: a.StartTime,
				EndTime:   a.EndTime,
			})
		}

		responseProviders = append(responseProviders, &pb.ListProviderR{
			UserId:      provider.UserId,
			CompanyName: provider.CompanyName,
			ServiceId:   serviceIds,
			Location: &pb.Location{
				City:    provider.Location.City,
				Country: provider.Location.Country,
			},
			Availabilitys: availabilitys,
			XId:           provider.XId.Hex(),
			CreratedAt:    provider.CreatedAt,
		})
	}

	resp := &pb.ListProvidersResponse{
		Listpriders: responseProviders,
	}

	return resp, nil
}

func (p *BookingRepo) SearchProviders(ctx context.Context, req *pb.SearchProvidersRequest) (*pb.SearchProvidersResponse, error) {
	collection := p.db.Collection("providers")

	filter := bson.M{}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}

	if req.CompanyName != "" {
		filter["company_name"] = bson.M{"$regex": req.CompanyName, "$options": "i"} 
	}

	if req.Location != nil {
		if req.Location.City != "" {
			filter["location.city"] = req.Location.City
		}
		if req.Location.Country != "" {
			filter["location.country"] = req.Location.Country
		}
	}

	if len(req.Availabilitys) > 0 {
		var availabilitys []bson.M
		for _, availability := range req.Availabilitys {
			availabilitys = append(availabilitys, bson.M{
				"start_time": availability.StartTime,
				"end_time":   availability.EndTime,
			})
		}
		filter["availabilitys"] = bson.M{"$all": availabilitys}
	}
	fmt.Println(filter)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Providerlarni qidirishda xatolik: %v", err)
	}
	defer cursor.Close(ctx)

	var providers []*pb.SearchProviderR
	for cursor.Next(ctx) {
		var provider pb.SearchProviderR
		if err := cursor.Decode(&provider); err != nil {
			return nil, status.Errorf(codes.Internal, "Natijalarni qayta ishlashda xatolik: %v", err)
		}
		providers = append(providers, &provider)
	}

	return &pb.SearchProvidersResponse{
		Provider: providers,
	}, nil
}
