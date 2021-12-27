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
		logger.Sugar().Errorf("Add user to group error: %v", err)
		return &npool.AddGroupUsersResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetGroupUsers(ctx context.Context, in *npool.GetGroupUsersRequest) (*npool.GetGroupUsersResponse, error) {
	resp, err := applicationgroupuser.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get group users error: %v", err)
		return &npool.GetGroupUsersResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) RemoveGroupUsers(ctx context.Context, in *npool.RemoveGroupUsersRequest) (*npool.RemoveGroupUsersResponse, error) {
	resp, err := applicationgroupuser.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("remove group users error: %v", err)
		return &npool.RemoveGroupUsersResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserGroup(ctx context.Context, in *npool.GetUserGroupRequest) (*npool.GetUserGroupResponse, error) {
	resp, err := applicationgroupuser.GetUserGroup(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user group info error: %v", err)
		return &npool.GetUserGroupResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}
