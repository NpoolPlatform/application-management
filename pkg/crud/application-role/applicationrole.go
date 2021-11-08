package applicationrole

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationrole"
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
	creator, err := uuid.Parse(in.Request.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
	}

	exist, err := RoleNameExist(ctx, in.Request)
	if err != nil || exist {
		return nil, xerrors.Errorf("role name has already exist in this app")
	}

	info, err := db.Client().
		ApplicationRole.
		Create().
		SetAppID(in.Request.AppID).
		SetRoleName(in.Request.RoleName).
		SetCreator(creator).
		SetAnnotation(in.Request.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create role: %v", err)
	}

	return &npool.CreateRoleResponse{
		Info: dbRowToApplicationRole(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetRoleRequest) (*npool.GetRoleResponse, error) {
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
	roleID, err := uuid.Parse(in.Request.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	exist, err := RoleNameExist(ctx, in.Request)
	if err != nil || exist {
		return nil, xerrors.Errorf("role name has already exist in this app")
	}

	query, err := db.Client().
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.ID(roleID),
				applicationrole.AppID(in.Request.AppID),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get role: %v", err)
	}

	if len(query) == 0 {
		return nil, xerrors.Errorf("role is not exist.")
	}

	info, err := db.Client().
		ApplicationRole.
		UpdateOneID(roleID).
		SetRoleName(in.Request.RoleName).
		SetAnnotation(in.Request.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update role: %v", err)
	}

	return &npool.UpdateRoleResponse{
		Info: dbRowToApplicationRole(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteRoleRequest) (*npool.DeleteRoleResponse, error) {
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

func RoleNameExist(ctx context.Context, in *npool.RoleInfo) (bool, error) {
	info, err := db.Client().
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.RoleName(in.RoleName),
				applicationrole.AppID(in.AppID),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to query role: %v", err)
	}

	if len(info) != 0 {
		return true, nil
	}

	return false, nil
}
