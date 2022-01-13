package application

import (
	"context"

	"github.com/NpoolPlatform/application-management/pkg/crud/application"
	applicationuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
	npool "github.com/NpoolPlatform/message/npool/application"
	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateApplicationRequest) (*npool.CreateApplicationResponse, error) {
	resp, err := application.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	_, err = applicationuser.Create(ctx, &npool.AddUsersToApplicationRequest{
		AppID:   resp.Info.ID,
		UserIDs: []string{in.Info.ApplicationOwner},
	})
	if err != nil {
		return nil, xerrors.Errorf("fail to add application owenr into application: %v", err)
	}

	return resp, nil
}
