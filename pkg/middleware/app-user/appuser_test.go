package appuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	applicationgroup "github.com/NpoolPlatform/application-management/pkg/crud/application-group"
	applicationgroupuser "github.com/NpoolPlatform/application-management/pkg/crud/application-group-user"
	"github.com/NpoolPlatform/application-management/pkg/middleware/application"
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

func TestAppUserMiddleware(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appInfo := &npool.ApplicationInfo{
		ApplicationName:  "test-app" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	resp1, err := application.Create(context.Background(), &npool.CreateApplicationRequest{
		Info: appInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.Info.ID, uuid.UUID{})
		assert.Equal(t, resp1.Info.ApplicationName, appInfo.ApplicationName)
		assert.Equal(t, resp1.Info.ApplicationOwner, appInfo.ApplicationOwner)
		appInfo.ID = resp1.Info.ID
	}

	applicationGroupInfo := &npool.GroupInfo{
		AppID:      appInfo.ID,
		GroupName:  uuid.New().String(),
		GroupOwner: appInfo.ApplicationOwner,
	}
	resp2, err := applicationgroup.Create(context.Background(), &npool.CreateGroupRequest{
		Info: applicationGroupInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp2.Info.ID, uuid.UUID{})
		assert.Equal(t, resp2.Info.AppID, applicationGroupInfo.AppID)
		assert.Equal(t, resp2.Info.GroupName, applicationGroupInfo.GroupName)
		assert.Equal(t, resp2.Info.GroupOwner, applicationGroupInfo.GroupOwner)
		applicationGroupInfo.ID = resp2.Info.ID
	}

	_, err = applicationgroupuser.Create(context.Background(), &npool.AddGroupUsersRequest{
		AppID:   applicationGroupInfo.AppID,
		GroupID: applicationGroupInfo.ID,
		UserIDs: []string{appInfo.ApplicationOwner},
	})
	assert.Nil(t, err)

	resp3, err := GetApplicationUserDetail(context.Background(), &npool.GetApplicationUserDetailRequest{
		AppID:  appInfo.ID,
		UserID: appInfo.ApplicationOwner,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp3)
	}
}
