package roleuser

import (
	"context"

	applicationrole "github.com/NpoolPlatform/application-management/pkg/crud/application-role"
	applicationroleuser "github.com/NpoolPlatform/application-management/pkg/crud/application-role-user"
	npool "github.com/NpoolPlatform/message/npool/application"
	"golang.org/x/xerrors"
)

func GetUserRole(ctx context.Context, in *npool.GetUserRoleRequest) (*npool.GetUserRoleResponse, error) {
	resp, err := applicationroleuser.GetUserRole(ctx, in)
	if err != nil {
		return nil, err
	}

	response := []*npool.RoleInfo{}

	for _, info := range resp.Info.Infos {
		resp, err := applicationrole.Get(ctx, &npool.GetRoleRequest{
			AppID:  in.AppID,
			RoleID: info.ID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail to get role: %v", err)
		}
		response = append(response, resp.Info)
	}
	return &npool.GetUserRoleResponse{
		Info: &npool.UserRole{
			Infos:  response,
			UserID: in.UserID,
			AppID:  in.AppID,
		},
	}, nil
}
