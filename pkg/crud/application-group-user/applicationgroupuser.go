package applicationgroupuser

import (
	"context"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroupuser"
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

func groupUserExist(ctx context.Context, userID, groupID uuid.UUID, appID string) (int, error) {
	info, err := db.Client().
		ApplicationGroupUser.
		Query().
		Where(
			applicationgroupuser.And(
				applicationgroupuser.DeleteAt(0),
				applicationgroupuser.UserID(userID),
				applicationgroupuser.GroupID(groupID),
				applicationgroupuser.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return -1, xerrors.Errorf("fail to get group user by name: %v", err)
	}

	if len(info) == 0 {
		return 0, nil
	} else if len(info) == 1 {
		return 1, nil
	}

	return -1, nil
}

func Create(ctx context.Context, in *npool.AddGroupUsersRequest) (*npool.AddGroupUsersResponse, error) {
	groupID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
	}

	response := []*npool.GroupUserInfo{}

	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		exist, err := groupUserExist(ctx, userID, groupID, in.AppID)
		if err != nil || exist != 0 {
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
	groupID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
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
	groupID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, xerrors.Errorf("invalid group id: %v", err)
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
