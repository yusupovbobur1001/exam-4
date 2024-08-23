package model

type UpdateProfileRequest struct {
	NewFirstName   string `json:"new_first_name"`
	NewPhoneNumber string `json:"new_phone_number"`
	NewRole        string `json:"new_role"`
}

type UpdateBookingRequest struct {
	UserId     string  `bson:"user_id"`
	ProviderId string  `bson:"provider_id"`
	ServiceId  string  `bson:"service_id"`
	Status     string  `bson:"status"`
	TatolPrice float32 `bson:"tatol_price"`
}
