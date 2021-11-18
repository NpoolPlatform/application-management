package applicationgroup

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroup"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplicationGroup(row *ent.ApplicationGroup) *npool.GroupInfo {
	return &npool.GroupInfo{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		GroupName:  row.GroupName,
		GroupOwner: row.GroupOwner.String(),
		GroupLogo:  row.GroupLogo,
		Annotation: row.Annotation,
		CreateAT:   row.CreateAt,
		UpdateAT:   row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateGroupRequest) (*npool.CreateGroupResponse, error) {
	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	if existGroupName, err := exist.GroupName(ctx, in.Info.GroupName, in.Info.AppID); err != nil || existGroupName == -1 {
		return nil, xerrors.Errorf("group name has already exist")
	}

	id, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	groupOwner, err := uuid.Parse(in.Info.GroupOwner)
	if err != nil {
		return nil, xerrors.Errorf("invalid group owner id: %v", err)
	}

	info, err := db.Client().
		ApplicationGroup.
		Create().
		SetAppID(id).
		SetGroupName(in.Info.GroupName).
		SetGroupOwner(groupOwner).
		SetGroupLogo(in.Info.GroupLogo).
		SetAnnotation(in.Info.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create group: %v", err)
	}

	return &npool.CreateGroupResponse{
		Info: dbRowToApplicationGroup(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGroupRequest) (*npool.UpdateGroupResponse, error) {
	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	if existGroupName, err := exist.GroupName(ctx, in.Info.GroupName, in.Info.AppID); err != nil || existGroupName == -1 {
		return nil, xerrors.Errorf("group name has already exist")
	}

	groupID, err := uuid.Parse(in.Info.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
	}

	id, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	query, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.ID(groupID),
				applicationgroup.AppID(id),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query group: %v", err)
	}

	if query.DeleteAt != 0 {
		return nil, xerrors.Errorf("group has alreay been delete")
	}

	info, err := db.Client().
		ApplicationGroup.
		UpdateOneID(groupID).
		SetGroupName(in.Info.GroupName).
		SetGroupLogo(in.Info.GroupLogo).
		SetAnnotation(in.Info.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update group: %v", err)
	}

	return &npool.UpdateGroupResponse{
		Info: dbRowToApplicationGroup(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGroupRequest) (*npool.GetGroupResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	groupID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	info, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.ID(groupID),
				applicationgroup.AppID(id),
				applicationgroup.DeleteAt(0),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get group: %v", err)
	}

	return &npool.GetGroupResponse{
		Info: dbRowToApplicationGroup(info),
	}, nil
}

func GetGroupByOwner(ctx context.Context, in *npool.GetGroupByOwnerRequest) (*npool.GetGroupByOwnerResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	ownerID, err := uuid.Parse(in.Owner)
	if err != nil {
		return nil, xerrors.Errorf("invalid owner id: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	infos, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.GroupOwner(ownerID),
				applicationgroup.DeleteAt(0),
				applicationgroup.AppID(id),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query owner's group: %v", err)
	}

	response := []*npool.GroupInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationGroup(info))
	}
	return &npool.GetGroupByOwnerResponse{
		Info: &npool.OwnerGroup{
			Infos: response,
			AppID: in.AppID,
			Owner: in.Owner,
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetAllGroupsRequest) (*npool.GetAllGroupsResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	infos, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.DeleteAt(0),
				applicationgroup.AppID(id),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get all groups: %v", err)
	}

	response := []*npool.GroupInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationGroup(info))
	}

	return &npool.GetAllGroupsResponse{
		Infos: response,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteGroupRequest) (*npool.DeleteGroupResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	groupID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
	}

	_, err = db.Client().
		ApplicationGroup.
		UpdateOneID(groupID).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete group: %v", err)
	}
	return &npool.DeleteGroupResponse{
		Info: "delete group successfully",
	}, nil
}
