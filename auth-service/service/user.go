package service

import (
	pb "auth_service/genproto/user"
	"auth_service/model"
	"auth_service/storage/repo"
	"context"
	"fmt"
	"log/slog"
)

type AuthServiceI interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error)
	DeleteUserProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error)
	GetByIdProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error)
	GetAllProfile(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesResponse, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	Logout(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error)
}

type AuthService struct {
	pb.UnimplementedAuthServer
	UserRepo repo.StoageI
	Logger   *slog.Logger
}

func NewAuthService(userRepo repo.StoageI, logger *slog.Logger) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		Logger:   logger,
	}
}

func (a *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println(2)
	return a.UserRepo.Register(ctx, req)
}

func (a *AuthService) UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	return a.UserRepo.UpdateUserProfile(ctx, req)
}

func (a *AuthService) DeleteUserProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {
	return a.UserRepo.DeleteUserProfile(ctx, req)
}

func (a *AuthService) GetByIdProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	return a.UserRepo.GetByIdProfile(ctx, req)
}

func (a *AuthService) GetAllProfile(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesResponse, error) {
	return a.UserRepo.GetAllProfile(ctx, req)
}

func (a *AuthService) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	return a.UserRepo.Login(ctx, req)
}

func (a *AuthService) Logout(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error) {
	return a.UserRepo.Logout(ctx, req)
}
