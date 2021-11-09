// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	applicationgroupuser "github.com/NpoolPlatform/application-management/pkg/crud/application-group-user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddGroupUsers(ctx context.Context, in *npool.AddGroupUsersRequest) (*npool.AddGroupUsersResponse, error) {
	resp, err := applicationgroupuser.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create group error: %v", err)
		return &npool.AddGroupUsersResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGroupUsers(ctx context.Context, in *npool.GetGroupUsersRequest) (*npool.GetGroupUsersResponse, error) {
	resp, err := applicationgroupuser.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get group users error: %v", err)
		return &npool.GetGroupUsersResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) RemoveGroupUsers(ctx context.Context, in *npool.RemoveGroupUsersRequest) (*npool.RemoveGroupUsersResponse, error) {
	resp, err := applicationgroupuser.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("remove group users error: %v", err)
		return &npool.RemoveGroupUsersResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
