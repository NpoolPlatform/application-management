package applicationgroupuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/crud/application"
	applicationgroup "github.com/NpoolPlatform/application-management/pkg/crud/application-group"
	applicationuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
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

func TestApplicationGroupUserCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationInfo := &npool.ApplicationInfo{
		ApplicationName:  "test-group-user" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	respApp, err := application.Create(context.Background(), &npool.CreateApplicationRequest{
		Request: applicationInfo,
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

	respGroup, err := applicationgroup.Create(context.Background(), &npool.CreateGroupRequest{
		Request: applicationGroup,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, respGroup.Info.ID, uuid.UUID{})
		assert.Equal(t, respGroup.Info.AppID, applicationGroup.AppID)
		assert.Equal(t, respGroup.Info.GroupName, applicationGroup.GroupName)
		assert.Equal(t, respGroup.Info.GroupOwner, applicationGroup.GroupOwner)
		applicationGroup.ID = respGroup.Info.ID
	}

	applicationUser := &npool.ApplicationUserInfo{
		AppID:    applicationInfo.ID,
		UserID:   uuid.New().String(),
		Original: true,
	}

	respUser, err := applicationuser.Create(context.Background(), &npool.AddUsersToApplicationRequest{
		AppID:    applicationUser.AppID,
		UserIDs:  []string{applicationUser.UserID},
		Original: applicationUser.Original,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, respUser.Infos[0].ID, uuid.UUID{})
		assert.Equal(t, respUser.Infos[0].AppID, applicationUser.AppID)
		assert.Equal(t, respUser.Infos[0].UserID, applicationUser.UserID)
		assert.Equal(t, respUser.Infos[0].Original, applicationUser.Original)
		applicationUser.ID = respUser.Infos[0].ID
	}

	groupUser := &npool.GroupUserInfo{
		AppID:   applicationInfo.ID,
		GroupID: applicationGroup.ID,
		UserID:  applicationUser.UserID,
	}

	resp, err := Create(context.Background(), &npool.AddGroupUsersRequest{
		UserIDs: []string{groupUser.UserID},
		AppID:   groupUser.AppID,
		GroupID: groupUser.GroupID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Infos[0].ID, uuid.UUID{})
		assert.Equal(t, resp.Infos[0].AppID, groupUser.AppID)
		assert.Equal(t, resp.Infos[0].UserID, groupUser.UserID)
		assert.Equal(t, resp.Infos[0].GroupID, groupUser.GroupID)
		groupUser.ID = resp.Infos[0].ID
	}

	resp1, err := Get(context.Background(), &npool.GetGroupUsersRequest{
		GroupID: groupUser.GroupID,
		AppID:   groupUser.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("get all group users resp1 is: %v", resp1)

	_, err = Delete(context.Background(), &npool.RemoveGroupUsersRequest{
		GroupID: groupUser.GroupID,
		AppID:   groupUser.AppID,
		UserIDs: []string{groupUser.UserID},
	})
	assert.Nil(t, err)
}
