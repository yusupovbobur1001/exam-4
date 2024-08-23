package mongodb

import (
	"context"
	"fmt"
	"time"

	pb "booking_service/genproto/booking"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p *BookingRepo) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	collection := p.db.Collection("reviews")

	review := bson.M{
		"booking_id":  req.GetBookingId(),
		"user_id":     req.GetUserId(),
		"provider_id": req.GetProcviderId(),
		"rating":      req.GetRating(),
		"comment":     req.GetComment(),
		"created_at":  time.Now().Format(time.RFC3339),
	}

	_, err := collection.InsertOne(ctx, review)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateReviewResponse{
		BookingId:   req.GetBookingId(),
		UserId:      req.GetUserId(),
		ProcviderId: req.GetProcviderId(),
		Rating:      req.GetRating(),
		Comment:     req.GetComment(),
		CreatedAt:   review["created_at"].(string),
	}

	return resp, nil
}

func (p *BookingRepo) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error) {
	collection := p.db.Collection("reviews")
	fmt.Println(req.UserId)
	filter := bson.M{"_id": req.GetXId()}
	fmt.Println(filter, "----------------------------")
	update := bson.M{
		"$set": bson.M{
			"user_id":     req.GetUserId(),
			"provider_id": req.GetProcviderId(),
			"rating":      req.GetRating(),
			"comment":     req.GetComment(),
			"updated_at":  time.Now().Format(time.RFC3339), // Save the update time
		},
	}
	
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateReviewResponse{
		UserId:      req.GetUserId(),
		ProcviderId: req.GetProcviderId(),
		Rating:      req.GetRating(),
		Comment:     req.GetComment(),
		UpdatedAt:   update["$set"].(bson.M)["updated_at"].(string),
	}

	return resp, nil
}

func (p *BookingRepo) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	collection := p.db.Collection("reviews")

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
		return nil, fmt.Errorf("no review found with the given ID")
	}

	resp := &pb.DeleteReviewResponse{
		Message: "Review successfully deleted",
	}

	return resp, nil
}

func (p *BookingRepo) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	collection := p.db.Collection("reviews")

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

	var reviews []struct {
		BookingId  string  `bson:"booking_id"`
		UserId     string  `bson:"user_id"`
		ProviderId string  `bson:"provider_id"`
		Rating     float64 `bson:"rating"`
		Comment    string  `bson:"comment"`
		XId        string  `bson:"_id"`
		CreatedAt  string  `bson:"created_at"`
	}

	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	var responseReviews []*pb.ListReviewR
	for _, review := range reviews {
		responseReviews = append(responseReviews, &pb.ListReviewR{
			BookingId:   review.BookingId,
			UserId:      review.UserId,
			ProcviderId: review.ProviderId,
			Rating:      float32(review.Rating),
			Comment:     review.Comment,
			XId:         review.XId,
			CreatedAt:   review.CreatedAt,
		})
	}

	resp := &pb.ListReviewsResponse{
		Listreviews: responseReviews,
	}

	return resp, nil
}
