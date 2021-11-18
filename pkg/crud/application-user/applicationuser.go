package applicationuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationuser"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/NpoolPlatform/application-management/pkg/rollback"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplicationUser(row *ent.ApplicationUser) *npool.ApplicationUserInfo {
	return &npool.ApplicationUserInfo{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		UserID:   row.UserID.String(),
		Original: row.Original,
		CreateAT: row.CreateAt,
	}
}

func genCreate(ctx context.Context, client *ent.Client, in *npool.AddUsersToApplicationRequest) ([]*npool.ApplicationUserInfo, error) {
	createResponse := []*npool.ApplicationUserInfo{}
	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		if existAppUser, err := exist.ApplicationUser(ctx, in.AppID, userID); err != nil || existAppUser {
			return nil, xerrors.Errorf("user has already exist in this app: %v", err)
		}

		info, err := client.
			ApplicationUser.
			Create().
			SetAppID(id).
			SetUserID(userID).
			SetOriginal(in.Original).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to add user to application: %v", err)
		}
		createResponse = append(createResponse, dbRowToApplicationUser(info))
	}
	return createResponse, nil
}

func Create(ctx context.Context, in *npool.AddUsersToApplicationRequest) (*npool.AddUsersToApplicationResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	response, err := rollback.WithTX(ctx, db.Client(), func(tx *ent.Tx) (interface{}, error) {
		return genCreate(ctx, tx.Client(), in)
	})
	if err != nil {
		return nil, err
	}

	return &npool.AddUsersToApplicationResponse{
		Infos: response.([]*npool.ApplicationUserInfo),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserFromApplicationRequest) (*npool.GetUserFromApplicationResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	info, err := db.Client().
		ApplicationUser.
		Query().
		Where(
			applicationuser.And(
				applicationuser.AppID(id),
				applicationuser.UserID(userID),
				applicationuser.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user: %v", err)
	}

	if len(info) == 0 {
		return nil, xerrors.Errorf("user is not exist")
	}
	return &npool.GetUserFromApplicationResponse{
		Info: dbRowToApplicationUser(info[0]),
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetUsersFromApplicationRequest) (*npool.GetUsersFromApplicationResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	infos, err := db.Client().
		ApplicationUser.
		Query().
		Where(
			applicationuser.And(
				applicationuser.AppID(id),
				applicationuser.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get all users from application:: %v", err)
	}

	getUsersResponse := []*npool.ApplicationUserInfo{}
	for _, info := range infos {
		getUsersResponse = append(getUsersResponse, dbRowToApplicationUser(info))
	}
	return &npool.GetUsersFromApplicationResponse{
		Infos: getUsersResponse,
	}, nil
}

func genDelete(ctx context.Context, client *ent.Client, in *npool.RemoveUsersFromApplicationRequest) (interface{}, error) {
	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		_, err = client.
			ApplicationUser.
			Update().
			Where(
				applicationuser.And(
					applicationuser.AppID(id),
					applicationuser.UserID(userID),
					applicationuser.DeleteAt(0),
				),
			).
			SetDeleteAt(uint32(time.Now().Unix())).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to remove user from applciation: %v", err)
		}
	}
	return nil, nil
}

func Delete(ctx context.Context, in *npool.RemoveUsersFromApplicationRequest) (*npool.RemoveUsersFromApplicationResponse, error) {
	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	if _, err := rollback.WithTX(ctx, db.Client(), func(tx *ent.Tx) (interface{}, error) {
		return genDelete(ctx, tx.Client(), in)
	}); err != nil {
		return nil, err
	}

	return &npool.RemoveUsersFromApplicationResponse{
		Info: "remove users from application successfully",
	}, nil
}
