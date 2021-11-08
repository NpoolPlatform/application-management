package roleuser

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	applicationrole "github.com/NpoolPlatform/application-management/pkg/crud/application-role"
	applicationroleuser "github.com/NpoolPlatform/application-management/pkg/crud/application-role-user"
	"golang.org/x/xerrors"
)

func GetUserRole(ctx context.Context, in *npool.GetUserRoleRequest) (*npool.GetUserRoleResponse, error) {
	resp, err := applicationroleuser.GetUserRole(ctx, in)
	if err != nil {
		return nil, err
	}

	response := []*npool.RoleInfo{}

	for _, info := range resp.Infos {
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
		Infos: response,
	}, nil
}
