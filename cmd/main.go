package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	httphandlers "github.com/VeeRomanoff/Lollipop/internal/app/http_handlers"
	"github.com/VeeRomanoff/Lollipop/internal/app/lollipop/api/lollipop_api"
	"github.com/VeeRomanoff/Lollipop/internal/database"
	lollipop "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
	"github.com/VeeRomanoff/Lollipop/internal/s3"
	"github.com/VeeRomanoff/Lollipop/internal/services/users_service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort    = ":7001"
	metricsPort = ":8080"
	gatewayPort = ":8090"
)

func main() {
	// TODO export sensitive variables to environment config
	minioConfig := s3.Config{
		Endpoint:  "localhost:9000",
		AccessKey: "minioadmin",
		SecretKey: "minioadmin",
		UseSSL:    false,
	}

	// Инициализация s3 хранилища
	client := s3.NewClient(minioConfig)

	// Инициализация http сервера в отдельной горутине дабы избежать блокировки ListenAndServe
	go func() {
		if err := initHTTPServer(client, ":8088"); err != nil {
			log.Fatal(err)
		}
	}()

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
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	conn, err := grpc.NewClient(
		"0.0.0.0:7001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}

	gatewayMux := runtime.NewServeMux()
	err = lollipop.RegisterLollipopHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		log.Fatalf("failed to register gateway handler: %v", err)
	}

	gatewayServer := &http.Server{
		Addr:    gatewayPort,
		Handler: gatewayMux,
	}

	go func() {
		fmt.Println("gateway server listening on " + gatewayPort)
		log.Fatalln(gatewayServer.ListenAndServe())
	}()

	lollipop.RegisterLollipopServer(grpcServer, service)
	grpc_prometheus.Register(grpcServer)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("Starting metrics on port: %s", metricsPort)
		if err := http.ListenAndServe(metricsPort, nil); err != nil {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on port: %s", grpcPort)

	if err = grpcServer.Serve(listener); err != nil {
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

func initHTTPServer(mediaStoreClient *s3.MinioStore, port string) error {
	httpHandler := httphandlers.NewHTTPHandler(mediaStoreClient)
	mux := http.NewServeMux()
	mux.HandleFunc("/upload-image", httpHandler.UploadImage)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	if err := http.ListenAndServe(port, mux); err != nil {
		return fmt.Errorf("Failed to listen HTTP server: %v", err)
	}

	return nil
}
