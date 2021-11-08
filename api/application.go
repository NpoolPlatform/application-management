package api

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/crud/application"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateApplication(ctx context.Context, in *npool.CreateApplicationRequest) (*npool.CreateApplicationResponse, error) {
	resp, err := application.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create application error: %v", err)
		return &npool.CreateApplicationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateApplication(ctx context.Context, in *npool.UpdateApplicationRequest) (*npool.UpdateApplicationResponse, error) {
	resp, err := application.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update application error: %v", err)
		return &npool.UpdateApplicationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetApplication(ctx context.Context, in *npool.GetApplicationRequest) (*npool.GetApplicationResponse, error) {
	resp, err := application.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get application error: %v", err)
		return &npool.GetApplicationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetApplications(ctx context.Context, in *npool.GetApplicationsRequest) (*npool.GetApplicationsResponse, error) {
	resp, err := application.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get applications error: %v", err)
		return &npool.GetApplicationsResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteApplication(ctx context.Context, in *npool.DeleteApplicationRequest) (*npool.DeleteApplicationResponse, error) {
	resp, err := application.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update application error: %v", err)
		return &npool.DeleteApplicationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
