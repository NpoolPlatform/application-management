// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"

	applicationroleuser "github.com/NpoolPlatform/application-management/pkg/crud/application-role-user"
	roleuser "github.com/NpoolPlatform/application-management/pkg/middleware/role-user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SetUserRole(ctx context.Context, in *npool.SetUserRoleRequest) (*npool.SetUserRoleResponse, error) {
	resp, err := applicationroleuser.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("set user role error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UnsetUserRole(ctx context.Context, in *npool.UnSetUserRoleRequest) (*npool.UnSetUserRoleResponse, error) {
	resp, err := applicationroleuser.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("unset user role error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRoleUsers(ctx context.Context, in *npool.GetRoleUsersRequest) (*npool.GetRoleUsersResponse, error) {
	resp, err := applicationroleuser.GetRoleUsers(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get role users error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserRole(ctx context.Context, in *npool.GetUserRoleRequest) (*npool.GetUserRoleResponse, error) {
	resp, err := roleuser.GetUserRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user's role error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UnSetUserRole(ctx context.Context, in *npool.UnSetUserRoleRequest) (*npool.UnSetUserRoleResponse, error) {
	resp, err := applicationroleuser.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("unset user role error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
