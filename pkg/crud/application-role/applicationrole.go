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
		AppID:      row.AppID.String(),
		RoleName:   row.RoleName,
		Creator:    row.Creator.String(),
		Annotation: row.Annotation,
		CreateAT:   row.CreateAt,
		UpdateAT:   row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateRoleRequest) (*npool.CreateRoleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}
	creator, err := uuid.Parse(in.Info.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
	}

	id, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		ApplicationRole.
		Create().
		SetAppID(id).
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	roleID, err := uuid.Parse(in.RoleID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.AppID(id),
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
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
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.AppID(id),
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	roleID, err := uuid.Parse(in.Info.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	id, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	query, err := cli.
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.ID(roleID),
				applicationrole.AppID(id),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get role: %v", err)
	}

	if query.DeleteAt != 0 {
		return nil, xerrors.Errorf("role has been already deleted.")
	}

	info, err := cli.
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	roleID, err := uuid.Parse(in.RoleID)
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	_, err = cli.
		ApplicationRole.
		UpdateOneID(roleID).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete role: %v", err)
	}

	return &npool.DeleteRoleResponse{
		Info: "delete role successfully",
	}, nil
}

func GetRoleByCreator(ctx context.Context, in *npool.GetRoleByCreatorRequest) (*npool.GetRoleByCreatorResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	creatorID, err := uuid.Parse(in.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
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
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.AppID(id),
				applicationrole.Creator(creatorID),
				applicationrole.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get role by creator: %v", err)
	}

	response := []*npool.RoleInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationRole(info))
	}

	return &npool.GetRoleByCreatorResponse{
		Info: &npool.CreatorRole{
			Infos:   response,
			AppID:   in.AppID,
			Creator: in.Creator,
		},
	}, nil
}
