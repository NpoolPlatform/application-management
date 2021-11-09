// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	applicationgroup "github.com/NpoolPlatform/application-management/pkg/crud/application-group"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGroup(ctx context.Context, in *npool.CreateGroupRequest) (*npool.CreateGroupResponse, error) {
	resp, err := applicationgroup.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create group error: %v", err)
		return &npool.CreateGroupResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGroup(ctx context.Context, in *npool.GetGroupRequest) (*npool.GetGroupResponse, error) {
	resp, err := applicationgroup.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get group error: %v", err)
		return &npool.GetGroupResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAllGroups(ctx context.Context, in *npool.GetAllGroupsRequest) (*npool.GetAllGroupsResponse, error) {
	resp, err := applicationgroup.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get all groups error: %v", err)
		return &npool.GetAllGroupsResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateGroup(ctx context.Context, in *npool.UpdateGroupRequest) (*npool.UpdateGroupResponse, error) {
	resp, err := applicationgroup.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update group error: %v", err)
		return &npool.UpdateGroupResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteGroup(ctx context.Context, in *npool.DeleteGroupRequest) (*npool.DeleteGroupResponse, error) {
	resp, err := applicationgroup.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete group error: %v", err)
		return &npool.DeleteGroupResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGroupByOwner(ctx context.Context, in *npool.GetGroupByOwnerRequest) (*npool.GetGroupByOwnerResponse, error) {
	resp, err := applicationgroup.GetGroupByOwner(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get group by owner error: %v", err)
		return &npool.GetGroupByOwnerResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
