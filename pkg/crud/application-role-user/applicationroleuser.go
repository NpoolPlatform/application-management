package applicationroleuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	approle "github.com/NpoolPlatform/application-management/pkg/crud/application-role"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/NpoolPlatform/application-management/pkg/rollback"
	"github.com/google/uuid"
	"golang.org/x/xerrors"

	crudappuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
)

func dbRowToApplication(row *ent.ApplicationRoleUser) *npool.RoleUserInfo {
	return &npool.RoleUserInfo{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		RoleID:   row.RoleID.String(),
		UserID:   row.UserID.String(),
		CreateAT: row.CreateAt,
	}
}

func preConditionJudge(ctx context.Context, roleIDString, appID string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, appID); err != nil || !existApp {
		return uuid.UUID{}, xerrors.Errorf("application does not exist: %v", err)
	}

	if _, err := approle.Get(ctx, &npool.GetRoleRequest{
		AppID:  appID,
		RoleID: roleIDString,
	}); err != nil {
		return uuid.UUID{}, xerrors.Errorf("role doesn't exist")
	}

	roleID, err := uuid.Parse(roleIDString)
	if err != nil {
		return uuid.UUID{}, xerrors.Errorf("invalid role id: %v", err)
	}

	return roleID, nil
}

func genCreate(ctx context.Context, client *ent.Client, roleID uuid.UUID, in *npool.SetUserRoleRequest) ([]*npool.RoleUserInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	response := []*npool.RoleUserInfo{}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		appID, err := uuid.Parse(in.AppID)
		if err != nil {
			return nil, xerrors.Errorf("invalid app id: %v", err)
		}

		if _, err := crudappuser.Get(ctx, &npool.GetUserFromApplicationRequest{
			AppID:  in.AppID,
			UserID: userIDString,
		}); err != nil {
			return nil, xerrors.Errorf("user does not exist: %v", err)
		}

		_, err = client.
			ApplicationRoleUser.
			Query().
			Where(
				applicationroleuser.And(
					applicationroleuser.UserID(userID),
					applicationroleuser.RoleID(roleID),
					applicationroleuser.AppID(appID),
					applicationroleuser.DeleteAt(0),
				),
			).All(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to get user role: %v", err)
		}

		info, err := client.
			ApplicationRoleUser.
			Create().
			SetRoleID(roleID).
			SetAppID(appID).
			SetUserID(userID).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to set role to user")
		}
		response = append(response, dbRowToApplication(info))
	}
	return response, nil
}

func Create(ctx context.Context, in *npool.SetUserRoleRequest) (*npool.SetUserRoleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	roleID, err := preConditionJudge(ctx, in.RoleID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	response, err := rollback.WithTX(ctx, cli, func(tx *ent.Tx) (interface{}, error) {
		return genCreate(ctx, tx.Client(), roleID, in)
	})
	if err != nil {
		return nil, err
	}

	return &npool.SetUserRoleResponse{
		Infos: response.([]*npool.RoleUserInfo),
	}, nil
}

func GetRoleUsers(ctx context.Context, in *npool.GetRoleUsersRequest) (*npool.GetRoleUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	roleID, err := preConditionJudge(ctx, in.RoleID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ApplicationRoleUser.
		Query().
		Where(
			applicationroleuser.And(
				applicationroleuser.AppID(id),
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

func GetUserRole(ctx context.Context, in *npool.GetUserRoleRequest) (*npool.GetUserRoleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	if _, err := crudappuser.Get(ctx, &npool.GetUserFromApplicationRequest{
		AppID:  in.AppID,
		UserID: in.UserID,
	}); err != nil {
		return nil, xerrors.Errorf("user does not exist: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ApplicationRoleUser.
		Query().
		Where(
			applicationroleuser.And(
				applicationroleuser.DeleteAt(0),
				applicationroleuser.AppID(id),
				applicationroleuser.UserID(userID),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user's roles: %v", err)
	}

	response := []*npool.RoleInfo{}
	for _, info := range infos {
		res := &npool.RoleInfo{}
		res.ID = dbRowToApplication(info).RoleID
		response = append(response, res)
	}

	return &npool.GetUserRoleResponse{
		Info: &npool.UserRole{
			Infos:  response,
			UserID: in.UserID,
			AppID:  in.AppID,
		},
	}, nil
}

func genDelete(ctx context.Context, client *ent.Client, roleID uuid.UUID, in *npool.UnSetUserRoleRequest) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		id, err := uuid.Parse(in.AppID)
		if err != nil {
			return nil, xerrors.Errorf("invalid app id: %v", err)
		}

		_, err = client.
			ApplicationRoleUser.
			Update().
			Where(
				applicationroleuser.And(
					applicationroleuser.DeleteAt(0),
					applicationroleuser.AppID(id),
					applicationroleuser.UserID(userID),
					applicationroleuser.RoleID(roleID),
				),
			).
			SetDeleteAt(uint32(time.Now().Unix())).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to unset user role: %v", err)
		}
	}
	return nil, nil
}

func Delete(ctx context.Context, in *npool.UnSetUserRoleRequest) (*npool.UnSetUserRoleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	roleID, err := preConditionJudge(ctx, in.RoleID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	if _, err = rollback.WithTX(ctx, cli, func(tx *ent.Tx) (interface{}, error) {
		return genDelete(ctx, tx.Client(), roleID, in)
	}); err != nil {
		return nil, err
	}

	return &npool.UnSetUserRoleResponse{
		Info: "unset user role successfully",
	}, nil
}
