package applicationgroupuser

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroupuser"
	"github.com/NpoolPlatform/application-management/pkg/exist"
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
	existApp, err := exist.Application(ctx, appID)
	if err != nil || !existApp {
		return uuid.UUID{}, xerrors.Errorf("application does not exist: %v", err)
	}

	groupID, err := uuid.Parse(groupIDString)
	if err != nil {
		return uuid.UUID{}, xerrors.Errorf("invalid group id: %v", err)
	}

	existGroup, err := exist.ApplicationGroup(ctx, groupID, appID)
	if err != nil || !existGroup {
		return uuid.UUID{}, xerrors.Errorf("group doesn't exist: %v", err)
	}
	return groupID, nil
}

func Create(ctx context.Context, in *npool.AddGroupUsersRequest) (*npool.AddGroupUsersResponse, error) {
	groupID, err := preConditionJudge(ctx, in.GroupID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	response := []*npool.GroupUserInfo{}

	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		existGroupUser, err := exist.GroupUser(ctx, userID, groupID, in.AppID)
		if err != nil || existGroupUser {
			return nil, xerrors.Errorf("user has already existed in this app group")
		}

		info, err := db.Client().
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
	return &npool.AddGroupUsersResponse{
		Infos: response,
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

func Delete(ctx context.Context, in *npool.RemoveGroupUsersRequest) (*npool.RemoveGroupUsersResponse, error) {
	groupID, err := preConditionJudge(ctx, in.GroupID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("pre condition not pass: %v", err)
	}

	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		_, err = db.Client().
			ApplicationGroupUser.
			Update().
			Where(
				applicationgroupuser.And(
					applicationgroupuser.UserID(userID),
					applicationgroupuser.GroupID(groupID),
					applicationgroupuser.AppID(in.AppID),
					applicationgroupuser.DeleteAt(0),
				),
			).Save(ctx)

		if err != nil {
			return nil, xerrors.Errorf("fail to remove group user: %v", err)
		}
	}
	return &npool.RemoveGroupUsersResponse{
		Info: "remove users from group successfully",
	}, nil
}
