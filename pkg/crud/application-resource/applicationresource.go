package applicationresource

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationresource"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplicationResource(row *ent.ApplicationResource) *npool.ResourceInfo {
	return &npool.ResourceInfo{
		ID:                  row.ID.String(),
		AppID:               row.AppID,
		ResourceName:        row.ResourceName,
		Type:                row.Type,
		ResourceDescription: row.ResourceDescription,
		Creator:             row.Creator.String(),
		CreateAT:            int32(row.CreateAt),
		UpdateAT:            int32(row.UpdateAt),
	}
}

func Create(ctx context.Context, in *npool.CreateResourceRequest) (*npool.CreateResourceResponse, error) {
	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	if existName, err := exist.ResourceName(ctx, in.Info.ResourceName, in.Info.AppID); err != nil || existName != 0 {
		return nil, xerrors.Errorf("resource has been exist: %v", err)
	}

	creator, err := uuid.Parse(in.Info.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
	}

	info, err := db.Client().
		ApplicationResource.
		Create().
		SetAppID(in.Info.AppID).
		SetResourceName(in.Info.ResourceName).
		SetType(in.Info.Type).
		SetResourceDescription(in.Info.ResourceDescription).
		SetCreator(creator).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create resource: %v", err)
	}
	return &npool.CreateResourceResponse{
		Info: dbRowToApplicationResource(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetResourceRequest) (*npool.GetResourceResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	resourceID, err := uuid.Parse(in.ResourceID)
	if err != nil {
		return nil, xerrors.Errorf("invalid resource id: %v", err)
	}

	info, err := db.Client().
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.ID(resourceID),
				applicationresource.AppID(in.AppID),
				applicationresource.DeleteAt(0),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query resource: %v", err)
	}
	return &npool.GetResourceResponse{
		Info: dbRowToApplicationResource(info),
	}, nil
}

func GetResourceByCreator(ctx context.Context, in *npool.GetResourceByCreatorRequest) (*npool.GetResourceByCreatorResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	creatorID, err := uuid.Parse(in.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
	}

	infos, err := db.Client().
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.Creator(creatorID),
				applicationresource.AppID(in.AppID),
				applicationresource.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query resource by creator: %v", err)
	}

	response := []*npool.ResourceInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationResource(info))
	}

	return &npool.GetResourceByCreatorResponse{
		Info: &npool.CreatorResource{
			Infos:   response,
			AppID:   in.AppID,
			Creator: in.Creator,
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetResourcesRequest) (*npool.GetResourcesResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	infos, err := db.Client().
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.AppID(in.AppID),
				applicationresource.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get resources of app: %v", err)
	}

	response := []*npool.ResourceInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationResource(info))
	}
	return &npool.GetResourcesResponse{
		Infos: response,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateResourceRequest) (*npool.UpdateResourceResponse, error) {
	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	if existName, err := exist.ResourceName(ctx, in.Info.ResourceName, in.Info.AppID); err != nil || existName == -1 {
		return nil, xerrors.Errorf("resource name has already exist")
	}

	resourceID, err := uuid.Parse(in.Info.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid resource id: %v", err)
	}

	query, err := db.Client().
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.ID(resourceID),
				applicationresource.AppID(in.Info.AppID),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query resource: %v", err)
	}

	if query.DeleteAt != 0 {
		return nil, xerrors.Errorf("resource has already been delete")
	}

	info, err := db.Client().
		ApplicationResource.
		UpdateOneID(resourceID).
		SetResourceName(in.Info.ResourceName).
		SetType(in.Info.Type).
		SetResourceDescription(in.Info.ResourceDescription).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update resource: %v", err)
	}
	return &npool.UpdateResourceResponse{
		Info: dbRowToApplicationResource(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteResourceRequest) (*npool.DeleteResourceResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	resourceID, err := uuid.Parse(in.ResourceID)
	if err != nil {
		return nil, xerrors.Errorf("invalid resource id: %v", err)
	}

	_, err = db.Client().
		ApplicationResource.
		UpdateOneID(resourceID).
		SetDeleteAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete resource: %v", err)
	}
	return &npool.DeleteResourceResponse{
		Info: "delete resource successfully",
	}, nil
}
