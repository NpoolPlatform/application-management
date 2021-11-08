package applicationroleuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/NpoolPlatform/application-management/pkg/exist"
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

func preConditionJudge(ctx context.Context, roleIDString, appID string) (uuid.UUID, error) {
	existApp, err := exist.Application(ctx, appID)
	if err != nil || !existApp {
		return uuid.UUID{}, xerrors.Errorf("application does not exist: %v", err)
	}

	roleID, err := uuid.Parse(roleIDString)
	if err != nil {
		return uuid.UUID{}, xerrors.Errorf("invalid role id: %v", err)
	}

	existRole, err := exist.ApplicationRole(ctx, roleID, appID)
	if err != nil || !existRole {
		return uuid.UUID{}, xerrors.Errorf("role doesn't exist")
	}
	return roleID, nil
}

func Create(ctx context.Context, in *npool.SetUserRoleRequest) (*npool.SetUserRoleResponse, error) {
	roleID, err := preConditionJudge(ctx, in.RoleID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	resposne := []*npool.RoleUserInfo{}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		existUser, err := exist.ApplicationUser(ctx, in.AppID, userID)
		if err != nil || !existUser {
			return nil, xerrors.Errorf("user doesn't exist")
		}

		has, err := exist.UserRole(ctx, userID, roleID, in.AppID)
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
	roleID, err := preConditionJudge(ctx, in.RoleID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
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
	roleID, err := preConditionJudge(ctx, in.RoleID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
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
