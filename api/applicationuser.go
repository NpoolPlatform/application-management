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
