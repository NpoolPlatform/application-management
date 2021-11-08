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
		AppID:      row.AppID,
		GroupName:  row.GroupName,
		GroupOwner: row.GroupOwner.String(),
		GroupLogo:  row.GroupLogo,
		Annotation: row.Annotation,
		CreateAT:   int32(row.CreateAt),
		UpdateAT:   int32(row.UpdateAt),
	}
}

func Create(ctx context.Context, in *npool.CreateGroupRequest) (*npool.CreateGroupResponse, error) {
	existApp, err := exist.Application(ctx, in.Request.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	existGroupName, err := exist.GroupName(ctx, in.Request.GroupName, in.Request.AppID)
	if err != nil || existGroupName != 0 {
		return nil, xerrors.Errorf("group already exist in this app: %v", err)
	}

	groupOwner, err := uuid.Parse(in.Request.GroupOwner)
	if err != nil {
		return nil, xerrors.Errorf("invalid group owner id: %v", err)
	}

	info, err := db.Client().
		ApplicationGroup.
		Create().
		SetAppID(in.Request.AppID).
		SetGroupName(in.Request.GroupName).
		SetGroupOwner(groupOwner).
		SetGroupLogo(in.Request.GroupLogo).
		SetAnnotation(in.Request.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create group: %v", err)
	}

	return &npool.CreateGroupResponse{
		Info: dbRowToApplicationGroup(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGroupRequest) (*npool.UpdateGroupResponse, error) {
	existApp, err := exist.Application(ctx, in.Request.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	existGroupName, err := exist.GroupName(ctx, in.Request.GroupName, in.Request.AppID)
	if err != nil || existGroupName == -1 {
		return nil, xerrors.Errorf("group name has already exist")
	}

	groupID, err := uuid.Parse(in.Request.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
	}

	query, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.ID(groupID),
				applicationgroup.AppID(in.Request.AppID),
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
		SetGroupName(in.Request.GroupName).
		SetGroupLogo(in.Request.GroupLogo).
		SetAnnotation(in.Request.Annotation).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update group: %v", err)
	}

	return &npool.UpdateGroupResponse{
		Info: dbRowToApplicationGroup(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGroupRequest) (*npool.GetGroupResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	groupID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
	}

	info, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.ID(groupID),
				applicationgroup.AppID(in.AppID),
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

func GetAll(ctx context.Context, in *npool.GetAllGroupsRequest) (*npool.GetAllGroupsResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	infos, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.DeleteAt(0),
				applicationgroup.AppID(in.AppID),
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
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	groupID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
	}

	_, err = db.Client().
		ApplicationGroup.
		UpdateOneID(groupID).
		SetDeleteAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete group: %v", err)
	}
	return &npool.DeleteGroupResponse{
		Info: "delete group successfully",
	}, nil
}
