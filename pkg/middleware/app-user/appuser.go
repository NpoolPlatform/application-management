package appuser

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	applicationgroupuser "github.com/NpoolPlatform/application-management/pkg/crud/application-group-user"
	applicationroleuser "github.com/NpoolPlatform/application-management/pkg/crud/application-role-user"
	applicationuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
	"golang.org/x/xerrors"
)

func GetApplicationUserDetail(ctx context.Context, in *npool.GetApplicationUserDetailRequest) (*npool.GetApplicationUserDetailResponse, error) {
	respUser, err := applicationuser.Get(ctx, &npool.GetUserFromApplicationRequest{
		AppID:  in.AppID,
		UserID: in.UserID,
	})
	if err != nil {
		return nil, xerrors.Errorf("get application user info error: %v", err)
	}

	applicationUserDetail := &npool.ApplicationUserDetail{}
	applicationUserDetail.UserApplicationInfo = respUser.Info

	respUserRole, err := applicationroleuser.GetUserRole(ctx, &npool.GetUserRoleRequest{
		AppID:  in.AppID,
		UserID: in.UserID,
	})
	if err != nil {
		applicationUserDetail.UserRoleInfo = &npool.UserRole{}
	} else {
		applicationUserDetail.UserRoleInfo = respUserRole.Info
	}

	respUserGroup, err := applicationgroupuser.GetUserGroup(ctx, &npool.GetUserGroupRequest{
		AppID:  in.AppID,
		UserID: in.UserID,
	})
	if err != nil {
		applicationUserDetail.UserGroupInfos = make([]*npool.GroupUserInfo, 0)
	} else {
		applicationUserDetail.UserGroupInfos = respUserGroup.Infos
	}

	return &npool.GetApplicationUserDetailResponse{
		Info: applicationUserDetail,
	}, nil
}
