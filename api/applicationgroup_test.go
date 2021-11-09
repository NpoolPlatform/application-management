package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestApplicationGroupAPI(t *testing.T) { // nolint
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
		}).Post("http://localhost:32759/v1/create/app")
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
		}).Post("http://localhost:32759/v1/create/group")
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

	response2 := npool.GetGroupResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGroupRequest{
			GroupID: groupInfo.ID,
			AppID:   groupInfo.AppID,
		}).Post("http://localhost:32759/v1/get/group")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())

		err := json.Unmarshal(resp2.Body(), &response2)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response2.Info.ID, uuid.UUID{})
			assert.Equal(t, response2.Info.GroupName, groupInfo.GroupName)
			assert.Equal(t, response2.Info.AppID, groupInfo.AppID)
			assert.Equal(t, response2.Info.GroupOwner, groupInfo.GroupOwner)
		}
	}

	response3 := npool.GetAllGroupsResponse{}
	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetAllGroupsRequest{
			AppID: groupInfo.AppID,
		}).Post("http://localhost:32759/v1/get/all/groups")
	if assert.Nil(t, err) {
		fmt.Println(err, resp3)
		assert.Equal(t, 200, resp3.StatusCode())
		err := json.Unmarshal(resp3.Body(), &response3)
		assert.Nil(t, err)
	}

	groupInfo.Annotation = "test group api"
	response4 := npool.UpdateGroupResponse{}
	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGroupRequest{
			Info: &groupInfo,
		}).Post("http://localhost:32759/v1/update/group")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())

		err := json.Unmarshal(resp4.Body(), &response4)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response4.Info.ID, uuid.UUID{})
			assert.Equal(t, response4.Info.GroupName, groupInfo.GroupName)
			assert.Equal(t, response4.Info.AppID, groupInfo.AppID)
			assert.Equal(t, response4.Info.GroupOwner, groupInfo.GroupOwner)
			assert.Equal(t, response4.Info.Annotation, groupInfo.Annotation)
		}
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteGroupRequest{
			AppID:   groupInfo.AppID,
			GroupID: groupInfo.ID,
		}).Post("http://localhost:32759/v1/delete/group")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
	}
}
