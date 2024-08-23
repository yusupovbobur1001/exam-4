package main

import (
	"auth_service/api"
	"auth_service/api/handler"
	"auth_service/config"
	"auth_service/genproto/user"
	"auth_service/logger"
	"auth_service/service"
	"auth_service/storage/postgres"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	cfg  := config.Load()
	log := logger.NewLogger()

	userRepo := postgres.NewAuthRepoManagement(db)
	userService := service.NewAuthService(userRepo, log)
	userHandler := handler.NewAuthenticaionHandlerImpl(userService, log)
	fmt.Println(userHandler)
	s := grpc.NewServer()
	user.RegisterAuthServer(s, userService)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", cfg.HTTP_PORT)
		if err != nil {
			fmt.Println(err)
			log.Error("...", "Error while listening on TCP: %v", err)
			return
		}

		fmt.Printf("gRPC server is listening on port %s\n", cfg.HTTP_PORT)

		if err := s.Serve(lis); err != nil {
			log.Error("service not listen", "Failed to serve: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("Gin server is listening on port %s\n", cfg.USER_SERVICE)
		auth := api.NewService(userHandler)
		router := auth.NewRouter()
		if err := router.Run(cfg.USER_SERVICE); err != nil {
			log.Error("Gin server failed to run: %v", "error", err.Error())
		}
	}()

	wg.Wait()
}
