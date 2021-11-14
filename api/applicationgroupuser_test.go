package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestApplicationGroupUserAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	application := npool.ApplicationInfo{
		ApplicationName:  "test-role" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	response := npool.CreateApplicationResponse{}
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateApplicationRequest{
			Info: &application,
		}).Post("http://localhost:50080/v1/create/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())

		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Info.ID, uuid.UUID{})
			assert.Equal(t, response.Info.ApplicationName, application.ApplicationName)
			assert.Equal(t, response.Info.ApplicationOwner, application.ApplicationOwner)
			application.ID = response.Info.ID
		}
	}

	groupInfo := npool.GroupInfo{
		GroupName:  "test-group-api" + uuid.New().String(),
		AppID:      application.ID,
		GroupOwner: uuid.New().String(),
	}

	response1 := npool.CreateGroupResponse{}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGroupRequest{
			Info: &groupInfo,
		}).Post("http://localhost:50080/v1/create/group")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())

		err := json.Unmarshal(resp1.Body(), &response1)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response1.Info.ID, uuid.UUID{})
			assert.Equal(t, response1.Info.GroupName, groupInfo.GroupName)
			assert.Equal(t, response1.Info.AppID, groupInfo.AppID)
			assert.Equal(t, response1.Info.GroupOwner, groupInfo.GroupOwner)
			groupInfo.ID = response1.Info.ID
		}
	}

	appUser := npool.ApplicationUserInfo{
		AppID:  application.ID,
		UserID: uuid.New().String(),
	}

	response2 := npool.AddUsersToApplicationResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AddUsersToApplicationRequest{
			AppID:   appUser.AppID,
			UserIDs: []string{appUser.UserID},
		}).Post("http://localhost:50080/v1/add/users/to/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		err := json.Unmarshal(resp2.Body(), &response2)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response2.Infos[0].ID, uuid.UUID{})
			assert.Equal(t, response2.Infos[0].UserID, appUser.UserID)
			assert.Equal(t, response2.Infos[0].AppID, appUser.AppID)
			appUser.ID = response2.Infos[0].ID
		}
	}

	groupUserInfo := npool.GroupUserInfo{
		UserID:  appUser.UserID,
		GroupID: groupInfo.ID,
		AppID:   application.ID,
	}

	response3 := npool.AddGroupUsersResponse{}
	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AddGroupUsersRequest{
			UserIDs: []string{groupUserInfo.UserID},
			AppID:   groupInfo.AppID,
			GroupID: groupUserInfo.GroupID,
		}).Post("http://localhost:50080/v1/add/group/users")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		err := json.Unmarshal(resp3.Body(), &response3)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response3.Infos[0].ID, uuid.UUID{})
			assert.Equal(t, response3.Infos[0].UserID, groupUserInfo.UserID)
			assert.Equal(t, response3.Infos[0].AppID, groupUserInfo.AppID)
			assert.Equal(t, response3.Infos[0].GroupID, groupUserInfo.GroupID)
		}
	}

	response4 := npool.GetGroupUsersResponse{}
	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGroupUsersRequest{
			AppID:   groupInfo.AppID,
			GroupID: groupUserInfo.GroupID,
		}).Post("http://localhost:50080/v1/get/group/users")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
		err := json.Unmarshal(resp4.Body(), &response4)
		if assert.Nil(t, err) {
			assert.Equal(t, response4.Infos[0].UserID, groupUserInfo.UserID)
			assert.Equal(t, response4.Infos[0].AppID, groupUserInfo.AppID)
			assert.Equal(t, response4.Infos[0].GroupID, groupUserInfo.GroupID)
		}
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.RemoveGroupUsersRequest{
			AppID:   groupUserInfo.AppID,
			GroupID: groupUserInfo.GroupID,
			UserIDs: []string{groupUserInfo.UserID},
		}).Post("http://localhost:50080/v1/remove/group/users")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
	}
}
