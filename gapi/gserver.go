package gapi

import (
	"auth-service/config"
	"auth-service/db"
	"auth-service/proto"
	"context"
	"log"
)

type ServerGRPC struct {
	config config.Config
	store  db.Store
	proto.UnimplementedAuthServiceServer
}

func NewGrpcServer(config config.Config, store *db.Store) ServerGRPC {
	return ServerGRPC{
		config: config,
		store:  *store,
	}
}

func (s *ServerGRPC) Login(ctx context.Context, in *proto.AuthRequest) (*proto.AuthResponse, error) {
	log.Println("gRPC called! Login request!")

	arg := db.LoginParam{
		Email:    in.Email,
		Password: in.Password,
	}

	// Execute query.
	result, err := s.store.AuthenticateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	// Create response object
	response := &proto.AuthResponse{
		Jwt: result.JWT,
	}

	return response, nil
}
