package api

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	applicationuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddUsersToApplication(ctx context.Context, in *npool.AddUsersToApplicationRequest) (*npool.AddUsersToApplicationResponse, error) {
	resp, err := applicationuser.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("add user to app error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserFromApplication(ctx context.Context, in *npool.GetUserFromApplicationRequest) (*npool.GetUserFromApplicationResponse, error) {
	resp, err := applicationuser.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user from app error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUsersFromApplication(ctx context.Context, in *npool.GetUsersFromApplicationRequest) (*npool.GetUsersFromApplicationResponse, error) {
	resp, err := applicationuser.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get users from app error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) RemoveUsersFromApplication(ctx context.Context, in *npool.RemoveUsersFromApplicationRequest) (*npool.RemoveUsersFromApplicationResponse, error) {
	resp, err := applicationuser.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("remove users from app error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
