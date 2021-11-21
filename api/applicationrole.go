// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	applicationrole "github.com/NpoolPlatform/application-management/pkg/crud/application-role"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRole(ctx context.Context, in *npool.CreateRoleRequest) (*npool.CreateRoleResponse, error) {
	resp, err := applicationrole.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create role error: %v", err)
		return &npool.CreateRoleResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateRole(ctx context.Context, in *npool.UpdateRoleRequest) (*npool.UpdateRoleResponse, error) {
	resp, err := applicationrole.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update role error: %v", err)
		return &npool.UpdateRoleResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetRole(ctx context.Context, in *npool.GetRoleRequest) (*npool.GetRoleResponse, error) {
	resp, err := applicationrole.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get role error: %v", err)
		return &npool.GetRoleResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetRoles(ctx context.Context, in *npool.GetRolesRequest) (*npool.GetRolesResponse, error) {
	resp, err := applicationrole.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update role error: %v", err)
		return &npool.GetRolesResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteRole(ctx context.Context, in *npool.DeleteRoleRequest) (*npool.DeleteRoleResponse, error) {
	resp, err := applicationrole.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete role error: %v", err)
		return &npool.DeleteRoleResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetRoleByCreator(ctx context.Context, in *npool.GetRoleByCreatorRequest) (*npool.GetRoleByCreatorResponse, error) {
	resp, err := applicationrole.GetRoleByCreator(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete role error: %v", err)
		return &npool.GetRoleByCreatorResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}
