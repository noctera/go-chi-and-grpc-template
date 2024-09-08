package grpcserver

import (
	"google.golang.org/grpc"
	"log"
	"my-app/internal/inbound/grpcserver"
	"my-app/internal/inbound/grpcserver/proto/pb/testendpoint"
	"my-app/internal/pkg/config"
	"my-app/internal/pkg/repository"
	"net"
)

func StartGRPCServer() {
	cfg := config.LoadConfig()
	// Initialize the database
	repository.Init(cfg.DatabaseURL)
	defer repository.Close()

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err)
	}

	grpcServer := grpc.NewServer()

	testendpoint.RegisterTestEndpointServer(grpcServer, &grpcserver.TestEndpointServer{})

	log.Println("Starting gRPC server on :9001")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
