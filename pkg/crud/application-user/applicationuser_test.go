package applicationuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/crud/application"
	testinit "github.com/NpoolPlatform/application-management/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestApplicationUserCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationInfo := &npool.ApplicationInfo{
		ApplicationName:  "test-user" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	respApp, err := application.Create(context.Background(), &npool.CreateApplicationRequest{
		Info: applicationInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, respApp.Info.ID, "")
		assert.NotEqual(t, respApp.Info.ClientSecret, "")
		assert.Equal(t, respApp.Info.ApplicationName, applicationInfo.ApplicationName)
		assert.Equal(t, respApp.Info.ApplicationOwner, applicationInfo.ApplicationOwner)
		applicationInfo.ID = respApp.Info.ID
	}

	applicationUser := &npool.ApplicationUserInfo{
		AppID:    applicationInfo.ID,
		UserID:   uuid.New().String(),
		Original: true,
	}

	resp, err := Create(context.Background(), &npool.AddUsersToApplicationRequest{
		AppID:    applicationUser.AppID,
		UserIDs:  []string{applicationUser.UserID},
		Original: applicationUser.Original,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Infos[0].ID, uuid.UUID{})
		assert.Equal(t, resp.Infos[0].AppID, applicationUser.AppID)
		assert.Equal(t, resp.Infos[0].UserID, applicationUser.UserID)
		assert.Equal(t, resp.Infos[0].Original, applicationUser.Original)
		applicationUser.ID = resp.Infos[0].ID
	}

	resp1, err := Get(context.Background(), &npool.GetUserFromApplicationRequest{
		AppID:  applicationUser.AppID,
		UserID: applicationUser.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, applicationUser.ID)
		assert.Equal(t, resp1.Info.UserID, applicationUser.UserID)
		assert.Equal(t, resp1.Info.AppID, applicationUser.AppID)
		assert.Equal(t, resp1.Info.Original, applicationUser.Original)
	}

	_, err = GetAll(context.Background(), &npool.GetUsersFromApplicationRequest{
		AppID: applicationUser.AppID,
	})
	assert.Nil(t, err)

	_, err = Delete(context.Background(), &npool.RemoveUsersFromApplicationRequest{
		UserIDs: []string{applicationUser.UserID},
		AppID:   applicationUser.AppID,
	})
	assert.Nil(t, err)
}
