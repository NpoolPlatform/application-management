// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/application-management/pkg/crud/application"
	middleware "github.com/NpoolPlatform/application-management/pkg/middleware/application"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateApplication(ctx context.Context, in *npool.CreateApplicationRequest) (*npool.CreateApplicationResponse, error) {
	if in.Info == nil || in.Info.ApplicationName == "" || in.Info.ApplicationOwner == "" {
		logger.Sugar().Errorf("invalid input params")
		return &npool.CreateApplicationResponse{}, status.Errorf(codes.InvalidArgument, "invalid input params")
	}

	resp, err := middleware.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create application error: %v", err)
		return &npool.CreateApplicationResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateApplication(ctx context.Context, in *npool.UpdateApplicationRequest) (*npool.UpdateApplicationResponse, error) {
	if in.Info == nil {
		logger.Sugar().Errorf("invalid input params")
		return &npool.UpdateApplicationResponse{}, status.Errorf(codes.InvalidArgument, "invalid input params")
	}
	resp, err := application.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update application error: %v", err)
		return &npool.UpdateApplicationResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetApplication(ctx context.Context, in *npool.GetApplicationRequest) (*npool.GetApplicationResponse, error) {
	resp, err := application.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get application error: %v", err)
		return &npool.GetApplicationResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetApplications(ctx context.Context, in *npool.GetApplicationsRequest) (*npool.GetApplicationsResponse, error) {
	resp, err := application.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get applications error: %v", err)
		return &npool.GetApplicationsResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteApplication(ctx context.Context, in *npool.DeleteApplicationRequest) (*npool.DeleteApplicationResponse, error) {
	resp, err := application.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update application error: %v", err)
		return &npool.DeleteApplicationResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetApplicationByOwner(ctx context.Context, in *npool.GetApplicationByOwnerRequest) (*npool.GetApplicationByOwnerResponse, error) {
	resp, err := application.GetApplicationByOwner(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get application by owner error: %v", err)
		return &npool.GetApplicationByOwnerResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}
