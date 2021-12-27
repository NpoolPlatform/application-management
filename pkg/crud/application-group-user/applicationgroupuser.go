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

	crudappuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
)

func dbRawToApplicationGroupUser(row *ent.ApplicationGroupUser) *npool.GroupUserInfo {
	return &npool.GroupUserInfo{
		UserID:     row.UserID.String(),
		AppID:      row.AppID.String(),
		GroupID:    row.GroupID.String(),
		Annotation: row.Annotation,
		CreateAT:   row.CreateAt,
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

	if _, err := Get(ctx, &npool.GetGroupUsersRequest{
		GroupID: groupIDString,
		AppID:   appID,
	}); err != nil {
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

		id, err := uuid.Parse(in.AppID)
		if err != nil {
			return nil, xerrors.Errorf("invalid app id: %v", err)
		}

		_, err = crudappuser.Get(ctx, &npool.GetUserFromApplicationRequest{
			AppID:  in.AppID,
			UserID: userID.String(),
		})
		if err != nil {
			return nil, xerrors.Errorf("user doesn't exist in this app.")
		}

		_, err = client.
			ApplicationGroupUser.
			Query().
			Where(
				applicationgroupuser.And(
					applicationgroupuser.GroupID(groupID),
					applicationgroupuser.AppID(id),
					applicationgroupuser.UserID(userID),
					applicationgroupuser.DeleteAt(0),
				),
			).All(ctx)

		if err != nil {
			return nil, xerrors.Errorf("user has already existed in this app group")
		}

		info, err := client.
			ApplicationGroupUser.
			Create().
			SetAppID(id).
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	response, err := rollback.WithTX(ctx, cli, func(tx *ent.Tx) (interface{}, error) {
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

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ApplicationGroupUser.
		Query().
		Where(
			applicationgroupuser.And(
				applicationgroupuser.DeleteAt(0),
				applicationgroupuser.GroupID(groupID),
				applicationgroupuser.AppID(id),
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

		id, err := uuid.Parse(in.AppID)
		if err != nil {
			return nil, xerrors.Errorf("invalid app id: %v", err)
		}

		_, err = client.
			ApplicationGroupUser.
			Update().
			Where(
				applicationgroupuser.And(
					applicationgroupuser.UserID(userID),
					applicationgroupuser.GroupID(groupID),
					applicationgroupuser.AppID(id),
					applicationgroupuser.DeleteAt(0),
				),
			).
			SetDeleteAt(uint32(time.Now().Unix())).
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	if _, err = rollback.WithTX(ctx, cli, func(tx *ent.Tx) (interface{}, error) {
		return genDelete(ctx, tx.Client(), groupID, in)
	}); err != nil {
		return nil, err
	}

	return &npool.RemoveGroupUsersResponse{
		Info: "remove users from group successfully",
	}, nil
}

func GetUserGroup(ctx context.Context, in *npool.GetUserGroupRequest) (*npool.GetUserGroupResponse, error) {
	appID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ApplicationGroupUser.
		Query().
		Where(
			applicationgroupuser.And(
				applicationgroupuser.DeleteAt(0),
				applicationgroupuser.UserID(userID),
				applicationgroupuser.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("get application user group error: %v", err)
	}

	response := []*npool.GroupUserInfo{}
	for _, info := range infos {
		response = append(response, dbRawToApplicationGroupUser(info))
	}

	return &npool.GetUserGroupResponse{
		Infos: response,
	}, nil
}
