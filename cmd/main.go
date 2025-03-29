package main

import (
	"github.com/VeeRomanoff/Lollipop/internal/app/lollipop/api/lollipop_api"
	"github.com/VeeRomanoff/Lollipop/internal/database"
	lollipop "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
	"github.com/VeeRomanoff/Lollipop/internal/services/users_service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":7001"

func main() {
	// Инициализация базы данных
	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Инициализация user service
	uService, err := initUserService(db)
	if err != nil {
		log.Fatalf("Failed to initialize user service: %v", err)
	}

	// Инициалищация gRPC сервиса
	service, err := initService(uService)
	if err != nil {
		log.Fatalf("Failed to initialize grpc service: %v", err)
	}
	// Инциализация gRPC сервера
	grpcServer := grpc.NewServer()
	log.Printf("Starting gRPC server...")

	lollipop.RegisterLollipopServer(grpcServer, service)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on port: %s", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initUserService(db *database.Database) (*users_service.Service, error) {
	return &users_service.Service{
		db,
	}, nil
}

func initService(userService *users_service.Service) (*lollipop_api.Implementation, error) {
	return lollipop_api.NewLollipop(
		userService,
	), nil
}
