package applicationuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationuser"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplicationUser(row *ent.ApplicationUser) *npool.ApplicationUserInfo {
	return &npool.ApplicationUserInfo{
		ID:       row.ID.String(),
		AppID:    row.AppID,
		UserID:   row.UserID.String(),
		Original: row.Original,
		CreateAT: int32(row.CreateAt),
	}
}

func Create(ctx context.Context, in *npool.AddUsersToApplicationRequest) (*npool.AddUsersToApplicationResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	createResponse := []*npool.ApplicationUserInfo{}
	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		existAppUser, err := exist.ApplicationUser(ctx, in.AppID, userID)
		if err != nil || existAppUser {
			return nil, xerrors.Errorf("user has already exist in this app: %v", err)
		}

		info, err := db.Client().
			ApplicationUser.
			Create().
			SetAppID(in.AppID).
			SetUserID(userID).
			SetOriginal(in.Original).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to add user to application: %v", err)
		}
		createResponse = append(createResponse, dbRowToApplicationUser(info))
	}
	return &npool.AddUsersToApplicationResponse{
		Infos: createResponse,
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserFromApplicationRequest) (*npool.GetUserFromApplicationResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	info, err := db.Client().
		ApplicationUser.
		Query().
		Where(
			applicationuser.And(
				applicationuser.AppID(in.AppID),
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

func GetAll(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	infos, err := db.Client().
		ApplicationUser.
		Query().
		Where(
			applicationuser.And(
				applicationuser.AppID(in.AppID),
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
	return &npool.GetUsersResponse{
		Infos: getUsersResponse,
	}, nil
}

func Delete(ctx context.Context, in *npool.RemoveUsersFromApplicationRequest) (*npool.RemoveUsersFromApplicationResponse, error) {
	existApp, err := exist.Application(ctx, in.AppID)
	if err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	for _, userIDString := range in.UserIDs {
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		_, err = db.Client().
			ApplicationUser.
			Update().
			Where(
				applicationuser.And(
					applicationuser.AppID(in.AppID),
					applicationuser.UserID(userID),
					applicationuser.DeleteAt(0),
				),
			).SetDeleteAt(time.Now().Unix()).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to remove user from applciation: %v", err)
		}
	}
	return &npool.RemoveUsersFromApplicationResponse{
		Info: "remove users from application successfully",
	}, nil
}
