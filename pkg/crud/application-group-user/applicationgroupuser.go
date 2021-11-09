package applicationgroupuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroupuser"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/NpoolPlatform/application-management/pkg/rollback"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRawToApplicationGroupUser(row *ent.ApplicationGroupUser) *npool.GroupUserInfo {
	return &npool.GroupUserInfo{
		UserID:     row.UserID.String(),
		AppID:      row.AppID,
		GroupID:    row.GroupID.String(),
		Annotation: row.Annotation,
		CreateAT:   int32(row.CreateAt),
	}
}

func preConditionJudge(ctx context.Context, groupIDString, appID string) (uuid.UUID, error) {
	if existApp, err := exist.Application(ctx, appID); err != nil || !existApp {
		return uuid.UUID{}, xerrors.Errorf("application does not exist: %v", err)
	}

	groupID, err := uuid.Parse(groupIDString)
	if err != nil {
		return uuid.UUID{}, xerrors.Errorf("invalid group id: %v", err)
	}

	if existGroup, err := exist.ApplicationGroup(ctx, groupID, appID); err != nil || !existGroup {
		return uuid.UUID{}, xerrors.Errorf("group doesn't exist: %v", err)
	}

	return groupID, nil
}

func genCreate(ctx context.Context, client *ent.Client, groupID uuid.UUID, in *npool.AddGroupUsersRequest) ([]*npool.GroupUserInfo, error) {
	response := []*npool.GroupUserInfo{}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		existUser, err := exist.ApplicationUser(ctx, in.AppID, userID)
		if err != nil || !existUser {
			return nil, xerrors.Errorf("user doesn't exist in this app.")
		}

		existGroupUser, err := exist.GroupUser(ctx, userID, groupID, in.AppID)
		if err != nil || existGroupUser {
			return nil, xerrors.Errorf("user has already existed in this app group")
		}

		info, err := client.
			ApplicationGroupUser.
			Create().
			SetAppID(in.AppID).
			SetUserID(userID).
			SetGroupID(groupID).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf(" fail to add user to app group: %v", err)
		}
		response = append(response, dbRawToApplicationGroupUser(info))
	}
	return response, nil
}

func Create(ctx context.Context, in *npool.AddGroupUsersRequest) (*npool.AddGroupUsersResponse, error) {
	groupID, err := preConditionJudge(ctx, in.GroupID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	response, err := rollback.WithTX(ctx, db.Client(), func(tx *ent.Tx) (interface{}, error) {
		return genCreate(ctx, tx.Client(), groupID, in)
	})
	if err != nil {
		return nil, err
	}

	return &npool.AddGroupUsersResponse{
		Infos: response.([]*npool.GroupUserInfo),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGroupUsersRequest) (*npool.GetGroupUsersResponse, error) {
	groupID, err := preConditionJudge(ctx, in.GroupID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	infos, err := db.Client().
		ApplicationGroupUser.
		Query().
		Where(
			applicationgroupuser.And(
				applicationgroupuser.DeleteAt(0),
				applicationgroupuser.GroupID(groupID),
				applicationgroupuser.AppID(in.AppID),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get all users from group: %v", err)
	}

	response := []*npool.GroupUserInfo{}
	for _, info := range infos {
		response = append(response, dbRawToApplicationGroupUser(info))
	}

	return &npool.GetGroupUsersResponse{
		Infos: response,
	}, nil
}

func genDelete(ctx context.Context, client *ent.Client, groupID uuid.UUID, in *npool.RemoveGroupUsersRequest) (interface{}, error) {
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		_, err = client.
			ApplicationGroupUser.
			Update().
			Where(
				applicationgroupuser.And(
					applicationgroupuser.UserID(userID),
					applicationgroupuser.GroupID(groupID),
					applicationgroupuser.AppID(in.AppID),
					applicationgroupuser.DeleteAt(0),
				),
			).
			SetDeleteAt(time.Now().Unix()).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to remove group user: %v", err)
		}
	}
	return nil, nil
}

func Delete(ctx context.Context, in *npool.RemoveGroupUsersRequest) (*npool.RemoveGroupUsersResponse, error) {
	groupID, err := preConditionJudge(ctx, in.GroupID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	if _, err = rollback.WithTX(ctx, db.Client(), func(tx *ent.Tx) (interface{}, error) {
		return genDelete(ctx, tx.Client(), groupID, in)
	}); err != nil {
		return nil, err
	}

	return &npool.RemoveGroupUsersResponse{
		Info: "remove users from group successfully",
	}, nil
}
