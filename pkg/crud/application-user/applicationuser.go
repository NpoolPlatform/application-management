package applicationuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationuser"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	"github.com/NpoolPlatform/application-management/pkg/rollback"
	npool "github.com/NpoolPlatform/message/npool/application"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplicationUser(row *ent.ApplicationUser) *npool.ApplicationUserInfo {
	return &npool.ApplicationUserInfo{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		UserID:      row.UserID.String(),
		KycVerify:   row.KycVerify,
		GAVerify:    row.GaVerify,
		GALogin:     row.GaLogin,
		SMSLogin:    row.SmsLogin,
		LoginNumber: row.LoginNumber,
		Original:    row.Original,
		CreateAT:    row.CreateAt,
	}
}

func genCreate(ctx context.Context, client *ent.Client, in *npool.AddUsersToApplicationRequest) ([]*npool.ApplicationUserInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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

		if _, err := Get(ctx, &npool.GetUserFromApplicationRequest{
			AppID:  in.AppID,
			UserID: userIDString,
		}); err == nil {
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	response, err := rollback.WithTX(ctx, cli, func(tx *ent.Tx) (interface{}, error) {
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	if _, err := rollback.WithTX(ctx, cli, func(tx *ent.Tx) (interface{}, error) {
		return genDelete(ctx, tx.Client(), in)
	}); err != nil {
		return nil, err
	}

	return &npool.RemoveUsersFromApplicationResponse{
		Info: "remove users from application successfully",
	}, nil
}

func SetSMSLogin(ctx context.Context, in *npool.SetSMSLoginRequest) (*npool.SetSMSLoginResponse, error) { // nolint
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

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

	_, err = cli.
		ApplicationUser.
		Update().
		Where(
			applicationuser.And(
				applicationuser.AppID(appID),
				applicationuser.UserID(userID),
				applicationuser.DeleteAt(0),
			),
		).
		SetSmsLogin(in.Set).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to set user sms login verify: %v", err)
	}
	return &npool.SetSMSLoginResponse{
		Info: "successfully set user sms login verify",
	}, nil
}

func SetGALogin(ctx context.Context, in *npool.SetGALoginRequest) (*npool.SetGALoginResponse, error) { // nolint
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	appID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	resp, err := Get(ctx, &npool.GetUserFromApplicationRequest{
		AppID:  in.AppID,
		UserID: in.UserID,
	})
	if err != nil {
		return nil, err
	}

	if !resp.Info.GAVerify {
		return nil, xerrors.Errorf("please bind your account to google authenticator")
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	_, err = cli.
		ApplicationUser.
		Update().
		Where(
			applicationuser.And(
				applicationuser.AppID(appID),
				applicationuser.UserID(userID),
				applicationuser.DeleteAt(0),
			),
		).
		SetGaLogin(in.Set).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to set user ga login: %v", err)
	}
	return &npool.SetGALoginResponse{
		Info: "successfully set user ga login verify",
	}, nil
}

func AddUserLoginTime(ctx context.Context, in *npool.AddUserLoginTimeRequest) (*npool.AddUserLoginTimeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

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

	info, err := cli.
		ApplicationUser.
		Query().
		Where(
			applicationuser.And(
				applicationuser.UserID(userID),
				applicationuser.AppID(appID),
				applicationuser.DeleteAt(0),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query app user: %v", err)
	}

	_, err = cli.
		ApplicationUser.
		UpdateOneID(info.ID).
		SetLoginNumber(info.LoginNumber + 1).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to add user login numbers: %v", err)
	}

	return &npool.AddUserLoginTimeResponse{
		Info: info.LoginNumber + 1,
	}, nil
}

func UpdateUserGAStatus(ctx context.Context, in *npool.UpdateUserGAStatusRequest) (*npool.UpdateUserGAStatusResponse, error) { // nolint
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	appID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	_, err = cli.
		ApplicationUser.
		Update().
		Where(
			applicationuser.And(
				applicationuser.DeleteAt(0),
				applicationuser.UserID(userID),
				applicationuser.AppID(appID),
			),
		).
		SetGaVerify(in.Status).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user ga status: %v", err)
	}

	return &npool.UpdateUserGAStatusResponse{
		Info: "successfully update user ga status: %v",
	}, nil
}

func UpdateUserKYCStatus(ctx context.Context, in *npool.UpdateUserKYCStatusRequest) (*npool.UpdateUserKYCStatusResponse, error) { // nolint
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	appID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	_, err = cli.
		ApplicationUser.
		Update().
		Where(
			applicationuser.And(
				applicationuser.DeleteAt(0),
				applicationuser.UserID(userID),
				applicationuser.AppID(appID),
			),
		).
		SetGaVerify(in.Status).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user kyc status: %v", err)
	}

	return &npool.UpdateUserKYCStatusResponse{
		Info: "successfully update user kyc status: %v",
	}, nil
}

func GetUserAppID(ctx context.Context, in *npool.GetUserAppIDRequest) (*npool.GetUserAppIDResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ApplicationUser.
		Query().
		Where(
			applicationuser.And(
				applicationuser.UserID(userID),
				applicationuser.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user app ids: %v", err)
	}

	response := []string{}
	for _, info := range infos {
		response = append(response, info.AppID.String())
	}
	return &npool.GetUserAppIDResponse{
		Infos: response,
	}, nil
}
