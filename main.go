package main

import (
	"log"
	"net"

	"auth-service/config"
	"auth-service/db"
	"auth-service/gapi"
	"auth-service/proto"

	"google.golang.org/grpc"
)

func main() {

	// Load configuration settings
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database
	store, err := db.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// GRPC server
	lis, err := net.Listen("tcp", config.GRPCAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := gapi.NewGrpcServer(config, store)
	grpcServer := grpc.NewServer()

	proto.RegisterAuthServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}

}
