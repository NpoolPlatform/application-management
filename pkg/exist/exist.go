package exist

import (
	"context"

	crudapp "github.com/NpoolPlatform/application-management/pkg/crud/application"
	npool "github.com/NpoolPlatform/message/npool/application"
	"golang.org/x/xerrors"
)

func Application(ctx context.Context, appID string) (bool, error) {
	_, err := crudapp.Get(ctx, &npool.GetApplicationRequest{
		AppID: appID,
	})
	if err != nil {
		return false, xerrors.Errorf("fail to get application: %v", err)
	}

	return true, nil
}
