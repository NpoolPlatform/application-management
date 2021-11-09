package applicationrole

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationrole"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplicationRole(row *ent.ApplicationRole) *npool.RoleInfo {
	return &npool.RoleInfo{
		ID:         row.ID.String(),
		AppID:      row.AppID,
		RoleName:   row.RoleName,
		Creator:    row.Creator.String(),
		Annotation: row.Annotation,
		CreateAT:   int32(row.CreateAt),
		UpdateAT:   int32(row.UpdateAt),
	}
}

func Create(ctx context.Context, in *npool.CreateRoleRequest) (*npool.CreateRoleResponse, error) {
	existApp, err := exist.Application(ctx, in.Info.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	creator, err := uuid.Parse(in.Info.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
	}

	existRoleName, err := exist.RoleName(ctx, in.Info.RoleName, in.Info.AppID)
	if err != nil || existRoleName != 0 {
		return nil, xerrors.Errorf("role name has already exist in this app")
	}

	info, err := db.Client().
		ApplicationRole.
		Create().
		SetAppID(in.Info.AppID).
		SetRoleName(in.Info.RoleName).
		SetCreator(creator).
		SetAnnotation(in.Info.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create role: %v", err)
	}

	return &npool.CreateRoleResponse{
		Info: dbRowToApplicationRole(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetRoleRequest) (*npool.GetRoleResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	roleID, err := uuid.Parse(in.RoleID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	info, err := db.Client().
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.AppID(in.AppID),
				applicationrole.ID(roleID),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get role from app: %v", err)
	}

	return &npool.GetRoleResponse{
		Info: dbRowToApplicationRole(info),
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetRolesRequest) (*npool.GetRolesResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	infos, err := db.Client().
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.AppID(in.AppID),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failt o get all roles from app: %v", err)
	}

	response := []*npool.RoleInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationRole(info))
	}
	return &npool.GetRolesResponse{
		Infos: response,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateRoleRequest) (*npool.UpdateRoleResponse, error) {
	existApp, err := exist.Application(ctx, in.Info.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	roleID, err := uuid.Parse(in.Info.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	existRoleName, err := exist.RoleName(ctx, in.Info.RoleName, in.Info.AppID)
	if err != nil || existRoleName == -1 {
		return nil, xerrors.Errorf("role name has already exist in this app")
	}

	query, err := db.Client().
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.ID(roleID),
				applicationrole.AppID(in.Info.AppID),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get role: %v", err)
	}

	if query.DeleteAt != 0 {
		return nil, xerrors.Errorf("role has been already deleted.")
	}

	info, err := db.Client().
		ApplicationRole.
		UpdateOneID(roleID).
		SetRoleName(in.Info.RoleName).
		SetAnnotation(in.Info.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update role: %v", err)
	}

	return &npool.UpdateRoleResponse{
		Info: dbRowToApplicationRole(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteRoleRequest) (*npool.DeleteRoleResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	roleID, err := uuid.Parse(in.RoleID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	_, err = db.Client().
		ApplicationRole.
		UpdateOneID(roleID).
		SetDeleteAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete role: %v", err)
	}

	return &npool.DeleteRoleResponse{
		Info: "delete role successfully",
	}, nil
}
