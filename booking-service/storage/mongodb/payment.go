package mongodb

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *BookingRepo) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	collection := p.db.Collection("payments")

	// To'lovni yaratish
	payment := bson.M{
		"booking_id":     req.GetBookingId(),
		"amount":         req.GetAmount(),
		"status":         req.GetStatus(),
		"payment_method": req.GetPaymentMethod(),
		"transaction_id": req.GetTransactionId(),
	}

	// To'lovni ma'lumotlar bazasiga saqlash
	resp, err := collection.InsertOne(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %v", err)
	}

	// ID'ni olish va formatlash
	oid, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID: %v", resp.InsertedID)
	}

	// Javobni tayyorlash
	return &pb.CreatePaymentResponse{
		BookingId:     req.GetBookingId(),
		Amount:        req.GetAmount(),
		Status:        req.GetStatus(),
		PaymentMethod: req.GetPaymentMethod(),
		TransactionId: req.GetTransactionId(),
		XId:           oid.Hex(),
	}, nil
}

func (p *BookingRepo) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	collection := p.db.Collection("payments")

	oid, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}

	var payment struct {
		BookingId     string  `bson:"booking_id"`
		Amount        float32 `bson:"amount"`
		Status        string  `bson:"status"`
		PaymentMethod string  `bson:"payment_method"`
		TransactionID string  `bson:"transaction_id"`
		XId           string  `bson:"_id"`
	}

	err = collection.FindOne(ctx, filter).Decode(&payment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "Payment with ID %s not found", req.GetXId())
		}
		return nil, err
	}

	resp := &pb.GetPaymentResponse{
		BookingId:     payment.BookingId,
		Amount:        payment.Amount,
		Status:        payment.Status,
		PaymentMethod: payment.PaymentMethod,
		TransactionId: payment.TransactionID,
		XId:           payment.XId,
	}

	return resp, nil
}

func (p *BookingRepo) ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error) {
	collection := p.db.Collection("payments")

	// Paginate parameters
	limit := int(req.GetLimit())
	offset := int(req.GetOffset())

	// MongoDB Find options
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	// Query payments
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var payments []struct {
		BookingId     string  `bson:"booking_id"`
		Amount        float32 `bson:"amount"`
		Status        string  `bson:"status"`
		PaymentMethod string  `bson:"payment_method"`
		TransactionID string  `bson:"transaction_id"`
		XId           string  `bson:"_id"`
	}

	if err := cursor.All(ctx, &payments); err != nil {
		return nil, err
	}

	// Convert to response format
	var responsePayments []*pb.ListPaymentR
	for _, payment := range payments {
		responsePayments = append(responsePayments, &pb.ListPaymentR{
			BookingId:     payment.BookingId,
			Amount:        payment.Amount,
			Status:        payment.Status,
			PaymentMethod: payment.PaymentMethod,
			TransactionId: payment.TransactionID,
			XId:         payment.XId,
		})
	}

	resp := &pb.ListPaymentsResponse{
		Getpayments: responsePayments,
	}

	return resp, nil
}
