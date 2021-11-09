package applicationgroup

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

func TestApplicationGroupCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationInfo := &npool.ApplicationInfo{
		ApplicationName:  "test-group" + uuid.New().String(),
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

	applicationGroup := &npool.GroupInfo{
		AppID:      applicationInfo.ID,
		GroupName:  "test" + uuid.New().String(),
		GroupOwner: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateGroupRequest{
		Info: applicationGroup,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.AppID, applicationGroup.AppID)
		assert.Equal(t, resp.Info.GroupName, applicationGroup.GroupName)
		assert.Equal(t, resp.Info.GroupOwner, applicationGroup.GroupOwner)
		applicationGroup.ID = resp.Info.ID
	}

	applicationGroup.GroupName = "test-update" + uuid.New().String()
	resp1, err := Update(context.Background(), &npool.UpdateGroupRequest{
		Info: applicationGroup,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, applicationGroup.ID)
		assert.Equal(t, resp1.Info.AppID, applicationGroup.AppID)
		assert.Equal(t, resp1.Info.GroupName, applicationGroup.GroupName)
		assert.Equal(t, resp1.Info.GroupOwner, applicationGroup.GroupOwner)
	}

	resp2, err := Get(context.Background(), &npool.GetGroupRequest{
		GroupID: applicationGroup.ID,
		AppID:   applicationGroup.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, applicationGroup.ID)
		assert.Equal(t, resp2.Info.AppID, applicationGroup.AppID)
		assert.Equal(t, resp2.Info.GroupName, applicationGroup.GroupName)
		assert.Equal(t, resp2.Info.GroupOwner, applicationGroup.GroupOwner)
	}

	resp3, err := GetAll(context.Background(), &npool.GetAllGroupsRequest{
		AppID: applicationGroup.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("get all group is resp3: %v", resp3)

	resp4, err := GetGroupByOwner(context.Background(), &npool.GetGroupByOwnerRequest{
		Owner: applicationGroup.GroupOwner,
		AppID: applicationGroup.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("get owner's application is resp4: %v", resp4)

	_, err = GetGroupByOwner(context.Background(), &npool.GetGroupByOwnerRequest{
		Owner: uuid.New().String(),
		AppID: uuid.New().String(),
	})
	assert.NotNil(t, err)

	_, err = Delete(context.Background(), &npool.DeleteGroupRequest{
		GroupID: applicationGroup.ID,
		AppID:   applicationGroup.AppID,
	})
	assert.Nil(t, err)
}
