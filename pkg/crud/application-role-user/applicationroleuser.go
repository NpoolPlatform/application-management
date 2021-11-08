package applicationroleuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplication(row *ent.ApplicationRoleUser) *npool.RoleUserInfo {
	return &npool.RoleUserInfo{
		ID:       row.ID.String(),
		AppID:    row.AppID,
		RoleID:   row.RoleID.String(),
		UserID:   row.UserID.String(),
		CreateAT: int32(row.CreateAt),
	}
}

func userHasRole(ctx context.Context, userID, roleID uuid.UUID, appID string) (bool, error) {
	info, err := db.Client().
		ApplicationRoleUser.
		Query().
		Where(
			applicationroleuser.And(
				applicationroleuser.DeleteAt(0),
				applicationroleuser.RoleID(roleID),
				applicationroleuser.UserID(userID),
				applicationroleuser.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return false, xerrors.Errorf("fail to query user has role: %v", err)
	}
	if len(info) != 0 {
		return true, nil
	}

	return false, nil
}

func Create(ctx context.Context, in *npool.SetUserRoleRequest) (*npool.SetUserRoleResponse, error) {
	roleID, err := uuid.Parse(in.RoleID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}
	resposne := []*npool.RoleUserInfo{}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		has, err := userHasRole(ctx, userID, roleID, in.AppID)
		if err != nil || has {
			return nil, xerrors.Errorf("user already has the role: %v", err)
		}

		info, err := db.Client().
			ApplicationRoleUser.
			Create().
			SetRoleID(roleID).
			SetAppID(in.AppID).
			SetUserID(userID).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to set role to user")
		}
		resposne = append(resposne, dbRowToApplication(info))
	}
	return &npool.SetUserRoleResponse{
		Infos: resposne,
	}, nil
}

func GetRoleUsers(ctx context.Context, in *npool.GetRoleUsersRequest) (*npool.GetRoleUsersResponse, error) {
	roleID, err := uuid.Parse(in.RoleID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	infos, err := db.Client().
		ApplicationRoleUser.
		Query().
		Where(
			applicationroleuser.And(
				applicationroleuser.AppID(in.AppID),
				applicationroleuser.RoleID(roleID),
				applicationroleuser.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get role's user: %v", err)
	}
	response := []*npool.RoleUserInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplication(info))
	}
	return &npool.GetRoleUsersResponse{
		Infos: response,
	}, nil
}

func Delete(ctx context.Context, in *npool.UnSetUserRoleRequest) (*npool.UnSetUserRoleResponse, error) {
	roleID, err := uuid.Parse(in.RoleID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		_, err = db.Client().
			ApplicationRoleUser.
			Update().
			Where(
				applicationroleuser.And(
					applicationroleuser.DeleteAt(0),
					applicationroleuser.AppID(in.AppID),
					applicationroleuser.UserID(userID),
					applicationroleuser.RoleID(roleID),
				),
			).SetDeleteAt(time.Now().Unix()).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to unset user role: %v", err)
		}
	}
	return &npool.UnSetUserRoleResponse{
		Info: "unset user role successfully",
	}, nil
}
