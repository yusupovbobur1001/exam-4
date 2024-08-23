package postgres

import (
	"context"
	"fmt"
	"testing"

	pb "auth_service/genproto/user"
	"auth_service/model"
)

func TestRegister(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	a := NewAuthRepoManagement(db)

	req := pb.RegisterRequest{
		FirstName:   "bo",
		LastName:    "bo;",
		Email:       "byo",
		Password:    "1001",
		PhoneNumber: "bo",
		Role:        "user",
	}

	resp, err := a.Register(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Log(resp)
}

func TestUpdateUserProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	a := NewAuthRepoManagement(db)

	req := pb.UpdateProfileRequest{
		NewFirstName:   "asdf",
		NewPhoneNumber: "asdfa",
		NewRole:        "admin",
		Id:             "7d0610dc-ed62-400e-95e8-d8d3e9907dd1",
	}

	resp, err := a.UpdateUserProfile(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Log(resp)
}

func TestDeleteUserProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	a := NewAuthRepoManagement(db)

	req := pb.DeleteProfileRequest{
		Id: "7d0610dc-ed62-400e-95e8-d8d3e9907dd1",
	}

	resp, err := a.DeleteUserProfile(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Log(resp)
}

func TestGetAllProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	a := NewAuthRepoManagement(db)

	req := pb.GetProfilesRequest{
		Limit:  1,
		Offset: 0,
	}

	resp, err := a.GetAllProfile(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(resp)
}

func TestGetByIdProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	a := NewAuthRepoManagement(db)

	req := pb.GetProfileRequest{
		Id: "cf4ba28a-fcaa-4d67-81a4-2bbafc1e2c26",
	}

	resp, err := a.GetByIdProfile(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Log(resp)
}

func TestLogin(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	a := NewAuthRepoManagement(db)

	req := model.LoginRequest{
		Email:    "byo",
		Password: "1001",
	}

	resp, er := a.Login(context.Background(), &req)
	if er != nil {
		fmt.Println(err)
	}

	t.Log(resp)
}
